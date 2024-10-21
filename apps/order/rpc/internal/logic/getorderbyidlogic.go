package logic

import (
	"context"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"webshop/apps/order/rpc/internal/svc"
	"webshop/apps/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderByIdLogic {
	return &GetOrderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderByIdLogic) GetOrderById(in *order.GetOrderByIdReq) (*order.GetOrderByIdResp, error) {
	// query order by id
	orderInfo, err := l.svcCtx.OrderModel.FindOne(l.ctx, strconv.FormatInt(in.Id, 10))
	if err != nil {
		l.Logger.Errorf("query order by id failed: %v", err)
		return nil, status.Errorf(codes.Internal, "query order by id failed")
	}

	// return order info
	return &order.GetOrderByIdResp{
		Order: &order.Orders{
			Id:          orderInfo.Id,
			Userid:      int64(orderInfo.Userid),
			Payment:     orderInfo.Payment,
			Paymenttype: orderInfo.Paymenttype,
			Shoppingid:  orderInfo.Shoppingid,
			Postage:     orderInfo.Postage,
			Status:      orderInfo.Status,
		},
	}, nil
}
