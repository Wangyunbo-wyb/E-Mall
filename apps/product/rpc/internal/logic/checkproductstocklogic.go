package logic

import (
	"context"

	"webshop/apps/product/rpc/internal/svc"
	"webshop/apps/product/rpc/product"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckProductStockLogic {
	return &CheckProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckProductStockLogic) CheckProductStock(in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	p, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.ProductId))
	if err != nil {
		l.Logger.Errorf("[CheckProductStock] failed to find product record, err: %s", err.Error())
		return nil, xerr.NewErrCode(xerr.ProductNotFound)
	}
	//check if stock is enough
	if p.Stock < in.Num {
		return nil, xerr.NewErrCode(xerr.StockNotEnough)
	}
	return &product.UpdateProductStockResponse{}, nil
}
