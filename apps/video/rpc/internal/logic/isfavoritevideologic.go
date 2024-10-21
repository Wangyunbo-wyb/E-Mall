package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsFavoriteVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsFavoriteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsFavoriteVideoLogic {
	return &IsFavoriteVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsFavoriteVideoLogic) IsFavoriteVideo(in *video.IsFavoriteVideoRequest) (*video.IsFavoriteVideoResponse, error) {
	// query cache to check if the record exists
	if l.svcCtx.Redis.
		Exists(l.ctx, genFavoriteVideoCacheKey(in.UserId, in.VideoId)).
		Val() == 1 {
		return &video.IsFavoriteVideoResponse{
			IsFavorite: true,
		}, nil
	}

	// query if the record exists in the database
	err := l.svcCtx.DB.
		Where("user_id = ? And video_id = ?", in.UserId, in.VideoId).
		First(&model.Favorite{}).Error

	// if the record does not exist
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return &video.IsFavoriteVideoResponse{
			IsFavorite: false,
		}, nil
	}

	if err != nil {
		l.Logger.Errorf("[IsFavoriteVideo] failed to query database, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoCacheErr)
	}

	// if the record exists, set the cache
	err = l.svcCtx.Redis.
		Set(l.ctx, genFavoriteVideoCacheKey(in.UserId, in.VideoId), 1, time.Hour).Err()
	if err != nil {
		l.Logger.Errorf("[IsFavoriteVideo] failed to set cache, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoCacheErr)
	}

	return &video.IsFavoriteVideoResponse{
		IsFavorite: true,
	}, nil
}

func genFavoriteVideoCacheKey(userId uint64, videoId int64) string {
	return fmt.Sprintf("biz#favorite#userId#videoId:%d:%d", userId, videoId)
}
