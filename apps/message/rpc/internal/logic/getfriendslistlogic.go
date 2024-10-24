package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"webshop/apps/message/rpc/contact"
	"webshop/apps/message/rpc/internal/model"
	"webshop/apps/message/rpc/internal/svc"
	"webshop/apps/message/rpc/service"
)

type GetFriendsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsListLogic {
	return &GetFriendsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendsListLogic) GetFriendsList(in *service.GetFriendsListRequest) (*service.GetFriendsListResponse, error) {
	var result []model.Friend

	err := l.svcCtx.DB.Where("user_id = ?", in.UserId).Select("friend_id").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return &contact.GetFriendsListResponse{
		FriendsId: func() []int64 {
			var friendsId []int64
			for _, v := range result {
				friendsId = append(friendsId, v.FriendId)
			}
			return friendsId
		}(),
	}, nil
}
