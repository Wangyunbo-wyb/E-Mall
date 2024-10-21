package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"
)

type DeleteVideoCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteVideoCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoCommentLogic {
	return &DeleteVideoCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteVideoCommentLogic) DeleteVideoComment(in *video.DeleteVideoCommentRequest) (*video.Empty, error) {
	if err := l.svcCtx.DB.
		Where("id = ?", in.CommentId).
		Delete(&model.Comment{}).Error; err != nil {
		l.Logger.Errorf("[DeleteVideoComment] failed to delete comment record, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoDBErr)
	}

	return &video.Empty{}, nil
}
