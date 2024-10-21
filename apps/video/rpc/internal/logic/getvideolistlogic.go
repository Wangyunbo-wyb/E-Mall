package logic

import (
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/threading"
	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListLogic {
	return &GetVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const CacheHotVideoList = "hot#video#list"

func (l *GetVideoListLogic) GetVideoList(in *video.GetVideoListRequest) (*video.GetVideoListResponse, error) {
	var videos []model.Video

	err := l.svcCtx.DB.
		Model(&model.Video{}).
		Where("created_at < ?", time.Unix(in.LatestTime, 0)).
		Order("created_at desc"). // sort by creation time in descending order, the newest is in the front
		Limit(int(in.Num)).
		Find(&videos).Error
	if err != nil {
		l.Logger.Errorf("[GetVideoList] failed to get video list, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoDBErr)
	}

	var videoList []*video.VideoInfo
	for _, v := range videos {
		// put the popular video into the cache
		if model.IsPopularVideo(v.FavoriteCount, v.CommentCount) {
			if result, err := l.svcCtx.Redis.Exists(l.ctx, genVideoInfoCacheKey(int64(v.ID))).Result(); result != HotVideo && err == nil {
				threading.GoSafe(func() {
					// if the cache does not exist, write to the cache
					if err := l.svcCtx.Redis.HSet(l.ctx, genVideoInfoCacheKey(int64(v.ID)), map[string]interface{}{
						"AuthorId":      v.AuthorId,
						"Title":         v.Title,
						"PlayUrl":       v.PlayUrl,
						"CoverUrl":      v.CoverUrl,
						"FavoriteCount": v.FavoriteCount,
						"CommentCount":  v.CommentCount,
						"CreatedAt":     v.CreatedAt.Unix(),
					}); err != nil {
						l.Logger.Errorf("[GetVideoList] failed to set the value into cache, err: %s", err)
					}
					// put the video into the hot video list
					err = l.svcCtx.Redis.LPush(l.ctx, CacheHotVideoList, v.ID).Err()
					if err != nil {
						l.Logger.Errorf("[GetVideoList] failed to push the value into the hot video list,err: %s", err)
					}
				})
			}
		}

		videoList = append(videoList, &video.VideoInfo{
			Id:            int64(v.ID),
			AuthorId:      v.AuthorId,
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			CreateTime:    v.CreatedAt.Unix(),
		})
	}

	return &video.GetVideoListResponse{
		VideoList: videoList,
	}, nil
}
