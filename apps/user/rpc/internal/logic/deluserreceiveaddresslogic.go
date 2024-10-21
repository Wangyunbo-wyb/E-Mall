package logic

import (
	"context"

	"github.com/pkg/errors"
	"webshop/apps/user/rpc/internal/model"
	"webshop/apps/user/rpc/internal/svc"
	"webshop/apps/user/rpc/user"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserReceiveAddressLogic {
	return &DelUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserReceiveAddressLogic) DelUserReceiveAddress(in *user.UserReceiveAddressDelReq) (*user.UserReceiveAddressDelRes, error) {
	_, err := l.svcCtx.UserReceiveAddressModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(xerr.NewErrMsg("数据不存在"), "收获地址不存在")
		}
		return nil, err
	}

	dbAddress := new(model.UserReceiveAddress)
	dbAddress.Id = in.GetId()
	dbAddress.IsDelete = 1
	err = l.svcCtx.UserReceiveAddressModel.UpdateIsDelete(l.ctx, dbAddress)
	if err != nil {
		l.Logger.Errorf("[DelUserReceiveAddress] Failed to delete user's ReceiveAddress err : %v , in :%+v", err, in)
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	return &user.UserReceiveAddressDelRes{}, nil
}
