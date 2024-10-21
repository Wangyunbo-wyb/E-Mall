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

type GetUserReceiveAddressListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserReceiveAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserReceiveAddressListLogic {
	return &GetUserReceiveAddressListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserReceiveAddressListLogic) GetUserReceiveAddressList(in *user.UserReceiveAddressListReq) (*user.UserReceiveAddressListRes, error) {
	receiveAddressesList, err := l.svcCtx.UserReceiveAddressModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		l.Logger.Errorf("[GetUserReceiveAddressList] Failed to get user's receive address list err : %v , in :%+v", err, in)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DbError), "Failed  get user's receive address list in :%+v", in)
	}
	var resp []*user.UserReceiveAddress
	for _, receiveAddresses := range receiveAddressesList {
		var pbAddress user.UserReceiveAddress
		_ = copier.Copy(&pbAddress, receiveAddresses)
		resp = append(resp, &pbAddress)
	}
	return &user.UserReceiveAddressListRes{
		List: resp,
	}, nil
}
