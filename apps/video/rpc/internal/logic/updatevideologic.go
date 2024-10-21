package logic

import (
	"context"

	"gorm.io/gorm/clause"
	"webshop/apps/video/rpc/internal/model"
	"webshop/apps/video/rpc/internal/svc"
	"webshop/apps/video/rpc/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVideoLogic {
	return &UpdateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVideoLogic) UpdateVideo(in *video.UpdateVideoRequest) (*video.Empty, error) {
	//start the transaction
	tx := l.svcCtx.DB.Begin()

	var newVideo model.Video
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", in.Video.Id).First(&newVideo).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	newVideo.CommentCount = in.Video.CommentCount
	newVideo.FavoriteCount = in.Video.FavoriteCount

	err = tx.Clauses(clause.Locking{Strength: "UPDATE"}).Save(&newVideo).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	//commit the transaction
	tx.Commit()

	return &video.Empty{}, nil
}
