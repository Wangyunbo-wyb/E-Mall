// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package orderclient

import (
	"context"

	"webshop/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddOrderReq         = order.AddOrderReq
	AddOrderResp        = order.AddOrderResp
	CreateOrderRequest  = order.CreateOrderRequest
	CreateOrderResponse = order.CreateOrderResponse
	GetOrderByIdReq     = order.GetOrderByIdReq
	GetOrderByIdResp    = order.GetOrderByIdResp
	Orderitem           = order.Orderitem
	Orders              = order.Orders
	OrdersRequest       = order.OrdersRequest
	OrdersResponse      = order.OrdersResponse

	Order interface {
		Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
		CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
		CreateOrderCheck(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
		RollbackOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
		CreateOrderDTM(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderResp, error)
		CreateOrderDTMRevert(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderResp, error)
		GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*GetOrderByIdResp, error)
	}

	defaultOrder struct {
		cli zrpc.Client
	}
)

func NewOrder(cli zrpc.Client) Order {
	return &defaultOrder{
		cli: cli,
	}
}

func (m *defaultOrder) Orders(ctx context.Context, in *OrdersRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.Orders(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrderCheck(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateOrderCheck(ctx, in, opts...)
}

func (m *defaultOrder) RollbackOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.RollbackOrder(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrderDTM(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderResp, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateOrderDTM(ctx, in, opts...)
}

func (m *defaultOrder) CreateOrderDTMRevert(ctx context.Context, in *AddOrderReq, opts ...grpc.CallOption) (*AddOrderResp, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.CreateOrderDTMRevert(ctx, in, opts...)
}

func (m *defaultOrder) GetOrderById(ctx context.Context, in *GetOrderByIdReq, opts ...grpc.CallOption) (*GetOrderByIdResp, error) {
	client := order.NewOrderClient(m.cli.Conn())
	return client.GetOrderById(ctx, in, opts...)
}
