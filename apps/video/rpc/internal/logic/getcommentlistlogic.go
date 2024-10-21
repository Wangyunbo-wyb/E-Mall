package logic

import (
	"context"

	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentListLogic) GetCommentList(in *video.GetCommentListRequest) (*video.GetCommentListResponse, error) {
	// get video comment list in reverse order
	var comments []model.Comment
	if err := l.svcCtx.DB.
		Where("video_id = ?", in.VideoId).
		Limit(model.PopularVideoStandard).
		Order("created_at").
		Find(&comments).Error; err != nil {
		l.Logger.Errorf("[GetCommentList] failed to get comment list, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoDBErr)
	}

	// encapsulate comment data
	commentList := make([]*video.Comment, 0, len(comments))
	for _, v := range comments {
		commentList = append(commentList, &video.Comment{
			Id:         int64(v.ID),
			AuthorId:   v.UserId,
			CreateTime: v.CreatedAt.Unix(),
			Content:    v.Content,
		})
	}

	return &video.GetCommentListResponse{
		CommentList: commentList,
	}, nil
}
