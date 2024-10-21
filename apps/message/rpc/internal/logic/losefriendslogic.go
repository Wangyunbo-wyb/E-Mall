package logic

import (
	"context"

	"Zhihu/application/message/rpc/contact"
	"Zhihu/application/message/rpc/internal/model"
	"Zhihu/application/message/rpc/internal/svc"
	"Zhihu/application/message/rpc/service"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoseFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoseFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoseFriendsLogic {
	return &LoseFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoseFriendsLogic) LoseFriends(in *service.LoseFriendsRequest) (*service.Empty, error) {
	tx := l.svcCtx.DB.Begin()

	if err := tx.Where("user_id = ? AND friend_id = ?", in.UserAId, in.UserBId).Delete(&model.Friend{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Where("user_id = ? AND friend_id = ?", in.UserBId, in.UserAId).Delete(&model.Friend{}).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &contact.Empty{}, nil
}
