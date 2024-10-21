package logic

import (
	"context"

	"webshop/apps/user/rpc/internal/model"
	"webshop/apps/user/rpc/internal/svc"
	"webshop/apps/user/rpc/user"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserCollectionLogic {
	return &AddUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserCollectionLogic) AddUserCollection(in *user.UserCollectionAddReq) (*user.UserCollectionAddRes, error) {
	dbCollection := new(model.UserCollection)
	dbCollection.Uid = uint64(in.Uid)
	dbCollection.ProductId = uint64(in.ProductId)
	_, err := l.svcCtx.UserCollectionModel.Insert(l.ctx, dbCollection)
	if err != nil {
		l.Logger.Errorf("[AddUserCollection] Failed to add user's Collection err : %v , in :%+v", err, in)
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	return &user.UserCollectionAddRes{}, nil
}
