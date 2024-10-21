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

type EditUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserReceiveAddressLogic {
	return &EditUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditUserReceiveAddressLogic) EditUserReceiveAddress(in *user.UserReceiveAddressEditReq) (*user.UserReceiveAddressEditRes, error) {
	//TODO: improve the code to add the transaction
	_, err := l.svcCtx.UserReceiveAddressModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, xerr.NewErrCode(xerr.UserAddressNotExist)
		}
		l.Logger.Errorf("[EditUserReceiveAddress] Failed to edit user's ReceiveAddress err : %v ", err)
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// if set default address, reset all default address
	if in.IsDefault == 1 {
		err := l.svcCtx.UserReceiveAddressModel.ResetDefaultAddress(l.ctx, in.Uid)
		if err != nil {
			l.Logger.Errorf("[EditUserReceiveAddress] Failed to edit user's ReceiveAddress err : %v , in :%+v", err, in)
			return nil, xerr.NewErrCode(xerr.DbError)
		}
	}
	dbAddress := new(model.UserReceiveAddress)
	errCopy := copier.Copy(&dbAddress, in)
	if errCopy != nil {
		return nil, errCopy
	}
	err = l.svcCtx.UserReceiveAddressModel.Update(l.ctx, dbAddress)
	if err != nil {
		l.Logger.Errorf("[EditUserReceiveAddress] Failed to edit user's ReceiveAddress err : %v , in :%+v", err, in)
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	return &user.UserReceiveAddressEditRes{}, nil
}
