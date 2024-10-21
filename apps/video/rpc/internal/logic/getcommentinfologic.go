package logic

import (
	"context"
	"errors"

	"gorm.io/gorm"
	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentInfoLogic {
	return &GetCommentInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentInfoLogic) GetCommentInfo(in *video.GetCommentInfoRequest) (*video.GetCommentInfoResponse, error) {
	var comment model.Comment
	err := l.svcCtx.DB.Where("id = ?", in.CommentId).First(&comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		l.Logger.Errorf("[GetCommentInfo] comment id %d not found", in.CommentId)
		return nil, xerr.NewErrCode(xerr.VideoCommentNotExist)
	}

	if err != nil {
		l.Logger.Errorf("[GetCommentInfo] failed to get comment info, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.VideoDBErr)
	}

	return &video.GetCommentInfoResponse{
		Id:          uint64(comment.ID),
		UserId:      comment.UserId,
		Content:     comment.Content,
		CreatedTime: comment.CreatedAt.Unix(),
	}, nil
}
