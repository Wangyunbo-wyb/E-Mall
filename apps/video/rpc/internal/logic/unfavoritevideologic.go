package logic

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"webshop/apps/video/mq"
	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFavoriteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFavoriteVideoLogic {
	return &UnFavoriteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnFavoriteVideoLogic) UnFavoriteVideo(in *video.UnFavoriteVideoRequest) (*video.Empty, error) {
	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// query the user's favorite record
		f := model.Favorite{}
		err := tx.
			Where("user_id = ? And video_id = ?", in.UserId, in.VideoId).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&f).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			l.Logger.Info("failed to get favorite record, record not found")
			return err
		}

		if err != nil {
			l.Logger.Errorf("[UnFavoriteVideo] failed to get favorite record because of the error of database, err: %s", err.Error())
			return err
		}

		// delete the record
		if err := tx.Where("user_id = ? And video_id = ?", in.UserId, in.VideoId).Delete(&f).Error; err != nil {
			return err
		}

		// the number of video favorites minus one
		if err := tx.Model(&model.Video{}).
			Where("id = ?", in.VideoId).
			Clauses(clause.Locking{Strength: "UPDATE"}).
			UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).
			Error; err != nil {
			l.Logger.Errorf("[UnFavoriteVideo] failed to update favorite count, err: %s", err.Error())
			return err
		}

		return nil
	})

	if err != nil {
		l.Logger.Errorf("[UnFavoriteVideo] the transaction failed, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoUnfavoriteErr)
	}

	// send a task to delete the cache
	task, err := mq.NewDelCacheTask(genFavoriteVideoCacheKey(in.UserId, in.VideoId))
	if err != nil {
		l.Logger.WithContext(l.ctx).Errorf("fail to create task: %v", err)
		return nil, xerr.NewErrCode(xerr.VideoMQErr)
	}
	if _, err := l.svcCtx.AsynqClient.Enqueue(task); err != nil {
		l.Logger.WithContext(l.ctx).Errorf("fail to send task: %v", err)
		return nil, xerr.NewErrCode(xerr.VideoMQErr)
	}

	return &video.Empty{}, nil
}
