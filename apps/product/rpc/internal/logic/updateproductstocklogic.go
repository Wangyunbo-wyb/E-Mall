package logic

import (
	"context"

	"webshop/apps/product/rpc/internal/svc"
	"webshop/apps/product/rpc/product"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStockLogic {
	return &UpdateProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductStockLogic) UpdateProductStock(in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	err := l.svcCtx.ProductModel.UpdateProductStock(l.ctx, in.ProductId, in.Num)
	if err != nil {
		l.Logger.Errorf("[UpdateProductStockLogic] failed to update product stock: %s", err)
		return nil, xerr.NewErrCode(xerr.DbError)
	}
	return &product.UpdateProductStockResponse{}, nil
}
