package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"webshop/apps/product/rpc/internal/svc"
	"webshop/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackProductStockLogic {
	return &RollbackProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RollbackProductStockLogic) RollbackProductStock(in *product.UpdateProductStockRequest) (*product.UpdateProductStockResponse, error) {
	err := l.svcCtx.ProductModel.UpdateProductStock(l.ctx, in.ProductId, -in.Num)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &product.UpdateProductStockResponse{}, nil
}
