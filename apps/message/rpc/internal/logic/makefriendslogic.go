package logic

import (
	"context"

	"Zhihu/application/message/rpc/contact"
	"Zhihu/application/message/rpc/internal/model"
	"Zhihu/application/message/rpc/internal/svc"
	"Zhihu/application/message/rpc/service"
	"github.com/zeromicro/go-zero/core/logx"
)

type MakeFriendsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeFriendsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeFriendsLogic {
	return &MakeFriendsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeFriendsLogic) MakeFriends(in *service.MakeFriendsRequest) (*service.Empty, error) {
	newFriendsA := model.Friend{
		UserId:   in.UserAId,
		FriendId: in.UserBId,
	}

	newFriendsB := model.Friend{
		UserId:   in.UserBId,
		FriendId: in.UserAId,
	}

	tx := l.svcCtx.DB.Begin()

	if err := tx.Create(&newFriendsA).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Create(&newFriendsB).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &contact.Empty{}, nil
}
