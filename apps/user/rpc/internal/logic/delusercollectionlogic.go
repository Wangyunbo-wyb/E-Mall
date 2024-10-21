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

type DelUserCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserCollectionLogic {
	return &DelUserCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserCollectionLogic) DelUserCollection(in *user.UserCollectionDelReq) (*user.UserCollectionDelRes, error) {
	_, err := l.svcCtx.UserCollectionModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.Wrap(xerr.NewErrMsg("数据不存在"), "该商品没有被收藏")
		}
		return nil, err
	}
	dbCollection := new(model.UserCollection)
	dbCollection.Id = uint64(in.Id)
	dbCollection.IsDelete = 1
	err = l.svcCtx.UserCollectionModel.UpdateIsDelete(l.ctx, dbCollection)
	if err != nil {
		l.Logger.Errorf("[DelUserCollection] Failed to delete user's Collection err : %v , in :%+v", err, in)
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	return &user.UserCollectionDelRes{}, nil
}
