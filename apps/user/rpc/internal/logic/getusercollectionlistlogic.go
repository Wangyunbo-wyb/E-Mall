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

type GetUserCollectionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCollectionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCollectionListLogic {
	return &GetUserCollectionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserCollectionListLogic) GetUserCollectionList(in *user.UserCollectionListReq) (*user.UserCollectionListRes, error) {
	collectionList, err := l.svcCtx.UserCollectionModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		l.Logger.Errorf("[GetUserCollectionList] Failed to get user's Collection list err : %v , in :%+v", err, in)
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	var resp []int64
	for _, collections := range collectionList {
		resp = append(resp, int64(collections.ProductId))
	}
	return &user.UserCollectionListRes{
		ProductId: resp,
	}, nil
}
