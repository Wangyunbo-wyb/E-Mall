package logic

import (
	"context"
	"fmt"

	"webshop/apps/pay/rpc/internal/model"
	"webshop/apps/pay/rpc/internal/svc"
	"webshop/apps/pay/rpc/pay"
	"webshop/pkg/snowflake"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	SN_PREFIX_THIRD_PAYMENT = "PMT" //第三方支付流水记录前缀
	SN_PREFIX_PRODUCT_ORDER = "HSO" //商品订单前缀
)

type CreatePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePaymentLogic {
	return &CreatePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePaymentLogic) CreatePayment(in *pay.CreatePaymentReq) (*pay.CreatePaymentResp, error) {
	Sn := fmt.Sprintf("%s_%s", SN_PREFIX_THIRD_PAYMENT, snowflake.GenIDString())

	data := new(model.ThirdPayment)
	data.UserId = in.UserId
	data.Sn = Sn
	data.PayMode = in.PayModel
	data.PayTotal = in.PayTotal
	data.OrderId = in.OrderSn
	data.ServiceType = in.ServiceType

	_, err := l.svcCtx.ThirdPaymentModel.Insert(l.ctx, data)
	if err != nil {
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	return &pay.CreatePaymentResp{
		Sn: data.Sn,
	}, nil
}
