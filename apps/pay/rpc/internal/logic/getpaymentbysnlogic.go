package logic

import (
	"context"
	"errors"
	"strconv"

	"webshop/apps/pay/rpc/internal/model"
	"webshop/apps/pay/rpc/internal/svc"
	"webshop/apps/pay/rpc/pay"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentBySnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentBySnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentBySnLogic {
	return &GetPaymentBySnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetPaymentBySn 根据sn查询流水记录
func (l *GetPaymentBySnLogic) GetPaymentBySn(in *pay.GetPaymentBySnReq) (*pay.GetPaymentBySnResp, error) {
	thirdPayment, err := l.svcCtx.ThirdPaymentModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		l.Logger.Errorf("[GetPaymentBySn] FindOneBySn db err: %v,sn : %s", err, in.Sn)
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	var resp pay.PaymentDetail
	if thirdPayment != nil {
		resp.OrderId = thirdPayment.OrderId
		resp.ServiceType = thirdPayment.ServiceType
		resp.UserId = thirdPayment.UserId
		resp.PayMode = thirdPayment.PayMode
		resp.TransactionId = thirdPayment.TransactionId
		resp.PayStatus = thirdPayment.PayStatus
		resp.PayTime = thirdPayment.PayTime.Unix()
		resp.PayTotal = thirdPayment.PayTotal
		resp.Sn = thirdPayment.Sn
		//TODO
		tradeState, _ := strconv.ParseInt(thirdPayment.TradeState, 10, 64)
		resp.TradeState = tradeState
		resp.CreateTime = thirdPayment.CreateTime.Unix()
		resp.PayTime = thirdPayment.PayTime.Unix()
	}

	return &pay.GetPaymentBySnResp{
		PaymentDetail: &resp,
	}, nil
}
