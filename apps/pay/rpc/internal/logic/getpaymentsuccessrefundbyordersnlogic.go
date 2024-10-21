package logic

import (
	"context"

	"webshop/apps/pay/rpc/internal/svc"
	"webshop/apps/pay/rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentSuccessRefundByOrderSnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentSuccessRefundByOrderSnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentSuccessRefundByOrderSnLogic {
	return &GetPaymentSuccessRefundByOrderSnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据订单sn查询流水记录
func (l *GetPaymentSuccessRefundByOrderSnLogic) GetPaymentSuccessRefundByOrderSn(in *pay.GetPaymentSuccessRefundByOrderSnReq) (*pay.GetPaymentSuccessRefundByOrderSnResp, error) {
	// todo: add your logic here and delete this line

	return &pay.GetPaymentSuccessRefundByOrderSnResp{}, nil
}
