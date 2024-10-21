package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"webshop/apps/user/rpc/internal/model"
	"webshop/apps/user/rpc/internal/svc"
	"webshop/apps/user/rpc/user"
	"webshop/pkg/xerr"
)

type AddUserReceiveAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserReceiveAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserReceiveAddressLogic {
	return &AddUserReceiveAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserReceiveAddressLogic) AddUserReceiveAddress(in *user.UserReceiveAddressAddReq) (*user.UserReceiveAddressAddRes, error) {
	dbAddress := new(model.UserReceiveAddress)
	dbAddress.Uid = uint64(in.GetUid())
	dbAddress.Name = in.GetName()
	dbAddress.Phone = in.GetPhone()
	dbAddress.Province = in.GetProvince()
	dbAddress.City = in.GetCity()
	dbAddress.IsDefault = uint64(in.GetIsDefault())
	dbAddress.PostCode = in.GetPostCode()
	dbAddress.Region = in.GetRegion()
	dbAddress.DetailAddress = in.GetDetailAddress()
	_, err := l.svcCtx.UserReceiveAddressModel.Insert(l.ctx, dbAddress)
	if err != nil {
		l.Logger.Errorf("[AddUserReceiveAddress] Failed to add user's ReceiveAddress err : %v , in :%+v", err, in)
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	return &user.UserReceiveAddressAddRes{}, nil
}
