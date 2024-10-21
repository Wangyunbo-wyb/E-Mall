package logic

import (
	"context"

	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	videoclient "webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteVideoListLogic {
	return &GetFavoriteVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFavoriteVideoListLogic) GetFavoriteVideoList(in *video.GetFavoriteVideoListRequest) (*video.GetFavoriteVideoListResponse, error) {
	// get user favorite video id (sort by time, from new to old)
	var favorites []model.Favorite

	//function preload is used to load the associated data of the model.By default, the associated data is loaded lazily, which means that the associated data is loaded when it is used for the first time.
	if err := l.svcCtx.DB.
		Where("user_id = ?", in.UserId).
		Preload("Video").
		Order("created_at desc").
		Find(&favorites).Error; err != nil {
		l.Logger.Errorf("[GetFavoriteVideoList] failed to get favorite video list, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoDBErr)
	}

	// encapsulate query results
	videoList := make([]*video.VideoInfo, 0, len(favorites))
	for _, v := range favorites {
		//it may exist dirty data, need to judge whether the video exists
		if v.Video.ID == 0 {
			//TODO: delete dirty data asynchronously
			//go func(videoID int64) {
			//	if err := l.svcCtx.DB.Delete(&model.Video{}, videoID).Error; err != nil {
			//		l.Logger.Errorf("[GetFavoriteVideoList] failed to delete dirty data, video id: %d, err: %s", videoID, err.Error())
			//	}
			//}(int64(v.ID))
			continue
		}

		videoInfo := &video.VideoInfo{
			Id:            int64(v.Video.ID),
			AuthorId:      v.Video.AuthorId,
			Title:         v.Video.Title,
			PlayUrl:       v.Video.PlayUrl,
			CoverUrl:      v.Video.CoverUrl,
			FavoriteCount: v.Video.FavoriteCount,
			CommentCount:  v.Video.CommentCount,
		}

		videoList = append(videoList, videoInfo)
	}

	return &videoclient.GetFavoriteVideoListResponse{
		VideoList: videoList,
	}, nil
}
