package logic

import (
	"context"
	"errors"

	"gorm.io/gorm/clause"
	"webshop/apps/video/mq"
	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type FavoriteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteVideoLogic {
	return &FavoriteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteVideoLogic) FavoriteVideo(in *video.FavoriteVideoRequest) (*video.Empty, error) {
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// query if the user has collected the video
		f := model.Favorite{}
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("user_id = ? And video_id = ?", in.UserId, in.VideoId).
			First(&f).Error

		// the collect record already exists
		if err == nil {
			return nil
		}

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// if the user has not liked the video, create a record
		f.VideoId = in.VideoId
		f.UserId = in.UserId
		if err := tx.Create(&f).Error; err != nil {
			return err
		}

		// the video is hot (in the cache), only update the cache and hand it over to the timed task to update the database
		result, err := l.svcCtx.Redis.Exists(l.ctx, genVideoInfoCacheKey(in.VideoId)).Result()
		if result == HotVideo {
			task, err := mq.NewAddCacheValueTask(genVideoInfoCacheKey(in.VideoId), "FavoriteCount", 1)
			if err != nil {
				return err
			}
			if _, err := l.svcCtx.AsynqClient.Enqueue(task); err != nil {
				return err
			}
		} else {
			if err != nil {
				return err
			}
			if err := tx.Model(&model.Video{}).
				Where("id = ?", in.VideoId).
				Clauses(clause.Locking{Strength: "UPDATE"}).
				UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).
				Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		l.Logger.Errorf("[FavoriteVideo] failed to create favorite record, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoFailToFavorite)
	}
	return &video.Empty{}, err
}
