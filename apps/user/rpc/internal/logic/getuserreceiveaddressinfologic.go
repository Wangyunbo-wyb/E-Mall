package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"webshop/apps/user/rpc/internal/model"
	"webshop/apps/user/rpc/internal/svc"
	"webshop/apps/user/rpc/user"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserReceiveAddressInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserReceiveAddressInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReceiveAddressInfoLogic {
	return &GetUserReceiveAddressInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserReceiveAddressInfoLogic) GetUserReceiveAddressInfo(in *user.UserReceiveAddressInfoReq) (*user.UserReceiveAddress, error) {
	receiveAddress, err := l.svcCtx.UserReceiveAddressModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(xerr.NewErrMsg("收获地址数据不存在"), "收获地址数据不存在")
		}
		return nil, err
	}
	var resp user.UserReceiveAddress
	_ = copier.Copy(&resp, receiveAddress)
	resp.CreateTime = receiveAddress.CreateTime.Unix()
	resp.UpdateTime = receiveAddress.UpdateTime.Unix()
	return &resp, nil
}
