package logic

import (
	"context"

	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListByAuthorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListByAuthorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListByAuthorLogic {
	return &GetVideoListByAuthorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVideoListByAuthorLogic) GetVideoListByAuthor(in *video.GetVideoListByAuthorRequest) (*video.GetVideoListByAuthorResponse, error) {
	//get author's video list (from new to old)
	var videos []model.Video
	err := l.svcCtx.DB.Where("author_id = ?", in.AuthorId).Order("created_at desc").Find(&videos).Error
	if err != nil {
		l.Logger.Errorf("[GetVideoListByAuthor] failed to get video list, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoDBErr)
	}

	videoList := make([]*video.VideoInfo, 0, len(videos))
	for _, v := range videos {
		videoInfo := &video.VideoInfo{
			Id:            int64(v.ID),
			AuthorId:      v.AuthorId,
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
		}
		videoList = append(videoList, videoInfo)
	}

	return &video.GetVideoListByAuthorResponse{
		VideoList: videoList,
	}, nil
}
