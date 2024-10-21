package logic

import (
	"context"

	"webshop/apps/product/rpc/internal/svc"
	"webshop/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationProductsLogic struct {
	ctx              context.Context
	svcCtx           *svc.ServiceContext
	productListLogic *ProductListLogic
	logx.Logger
}

func NewOperationProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationProductsLogic {
	return &OperationProductsLogic{
		ctx:              ctx,
		svcCtx:           svcCtx,
		Logger:           logx.WithContext(ctx),
		productListLogic: NewProductListLogic(ctx, svcCtx),
	}
}

const (
	validStatus          = 1
	operationProductsKey = "operation#products"
)

func (l *OperationProductsLogic) OperationProducts(in *product.OperationProductsRequest) (*product.OperationProductsResponse, error) {
	//get products from cache
	opProducts, ok := l.svcCtx.LocalCache.Get(operationProductsKey)
	if ok {
		return &product.OperationProductsResponse{Products: opProducts.([]*product.ProductItem)}, nil
	}

	pos, err := l.svcCtx.OperationModel.OperationProducts(l.ctx, validStatus)
	if err != nil {
		return nil, err
	}
	var pids []int64
	for _, p := range pos {
		pids = append(pids, int64(p.ProductId))
	}
	products, err := l.productListLogic.productsByIds(l.ctx, pids)
	if err != nil {
		return nil, err
	}
	var pItems []*product.ProductItem
	for _, p := range products {
		pItems = append(pItems, &product.ProductItem{
			ProductId: int64(p.Id),
			Name:      p.Name,
		})
	}
	// put the product list into cache
	l.svcCtx.LocalCache.Set(operationProductsKey, pItems)
	return &product.OperationProductsResponse{Products: pItems}, nil
}
