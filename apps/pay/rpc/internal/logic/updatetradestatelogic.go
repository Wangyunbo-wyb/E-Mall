package logic

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"webshop/apps/pay/rpc/internal/model"
	"webshop/apps/pay/rpc/internal/svc"
	"webshop/apps/pay/rpc/pay"
	"webshop/pkg/kafka"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTradeStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTradeStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTradeStateLogic {
	return &UpdateTradeStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UpdateTradeState 更新交易状态
func (l *UpdateTradeStateLogic) UpdateTradeState(in *pay.UpdateTradeStateReq) (*pay.UpdateTradeStateResp, error) {
	// confirm payment
	thirdPayment, err := l.svcCtx.ThirdPaymentModel.FindOneBySn(l.ctx, in.Sn)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		l.Logger.Errorf("[UpdateTradeState] FindOneBySn db err: %v,sn : %s", err, in.Sn)
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	if thirdPayment == nil {
		l.Logger.Errorf("[UpdateTradeState] third payment record no exists,sn : %s", in.Sn)
		return nil, xerr.NewErrCode(xerr.PaymentNotExist)
	}

	//Judge the status of the payment record
	if in.PayStatus == model.ThirdPaymentPayTradeStateSuccess || in.PayStatus == model.ThirdPaymentPayTradeStateFAIL {
		if thirdPayment.PayStatus == model.ThirdPaymentPayTradeStateWait {
			return &pay.UpdateTradeStateResp{}, nil
		}
	} else if in.PayStatus == model.ThirdPaymentPayTradeStateRefund {
		if thirdPayment.PayStatus != model.ThirdPaymentPayTradeStateSuccess {
			return nil, xerr.NewErrCode(xerr.PaymentFailToRefund)
		}
	} else {
		return nil, xerr.NewErrCode(xerr.PaymentStatusNotSupport)
	}

	//update the payment record
	thirdPayment.TradeState = in.TradeState
	thirdPayment.TransactionId = in.TransactionId
	thirdPayment.TradeType = in.TradeType
	thirdPayment.TradeStateDesc = in.TradeStateDesc
	thirdPayment.PayStatus = in.PayStatus
	thirdPayment.PayTime = time.Unix(in.PayTime, 0)

	if err := l.svcCtx.ThirdPaymentModel.UpdateWithVersion(l.ctx, nil, thirdPayment); err != nil {
		l.Logger.Errorf("[UpdateTradeState] Update db err: %v,sn : %s", err, in.Sn)
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	//send kafka message
	if err := l.pubKqPaySuccess(l.ctx, in.Sn, in.PayStatus); err != nil && !errors.Is(err, MarshalError) {
		l.Logger.Errorf("[UpdateTradeState] pubKqPaySuccess err: %v,sn : %s", err, in.Sn)
	}

	return &pay.UpdateTradeStateResp{}, nil
}

var MarshalError = errors.New("json.Marshal err")

// pubKqPaySuccess send kafka message
func (l *UpdateTradeStateLogic) pubKqPaySuccess(ctx context.Context, sn string, payStatus int64) error {

	m := kafka.ThirdPaymentUpdatePayStatusNotifyMessage{
		Sn:        sn,
		PayStatus: payStatus,
	}

	body, err := json.Marshal(m)
	if err != nil {
		l.Logger.Errorf("[pubKqPaySuccess] json.Marshal err: %v", err)
		return MarshalError
	}

	return l.svcCtx.KafkaPusher.Push(ctx, string(body))
}
