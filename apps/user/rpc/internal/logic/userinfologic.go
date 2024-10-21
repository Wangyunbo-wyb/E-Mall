package logic

import (
	"context"

	"github.com/pkg/errors"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/status"
	"webshop/apps/user/rpc/internal/model"
	"webshop/apps/user/rpc/internal/svc"
	"webshop/apps/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	umsMember, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	var resp user.UserInfo
	_ = copier.Copy(&resp, umsMember)
	resp.CreateTime = umsMember.CreateTime.Unix()
	resp.UpdateTime = umsMember.UpdateTime.Unix()
	return &user.UserInfoResponse{
		User: &resp,
	}, nil
}
