package logic

import (
	"context"
	"fmt"

	"webshop/apps/order/rpc/internal/svc"
	"webshop/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type RollbackOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRollbackOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RollbackOrderLogic {
	return &RollbackOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RollbackOrderLogic) RollbackOrder(in *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	fmt.Printf("RollbackOrder: %+v\n", in)
	return &order.CreateOrderResponse{}, nil
}
