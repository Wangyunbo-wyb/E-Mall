package logic

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"webshop/apps/order/rpc/internal/svc"
	"webshop/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OrdersLogic) Orders(in *order.OrdersRequest) (*order.OrdersResponse, error) {
	// 查询用户最新创建的订单
	resOrder, err := l.svcCtx.OrderModel.FindOneByUid(l.ctx, in.UserId)
	if err != nil {
		l.Logger.Errorf("failed to get the info of the orders: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get the info of the orders")
	}
	//TODO:
	orders := []*order.Orderitem{
		{
			Orderid: resOrder.Id,
			Userid:  int64(resOrder.Userid),
		},
	}
	return &order.OrdersResponse{Orders: orders}, nil // 返回订单列表
}
