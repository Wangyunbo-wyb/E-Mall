package model

import (
	"errors"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var ErrNotFound = sqlx.ErrNotFound
var ErrNoRowsUpdate = errors.New("update db no rows change")

// third party payment status
var ThirdPaymentPayTradeStateFAIL int64 = -1   //payment failed
var ThirdPaymentPayTradeStateWait int64 = 0    //waiting for payment
var ThirdPaymentPayTradeStateSuccess int64 = 1 //payment success
var ThirdPaymentPayTradeStateRefund int64 = 2  //refund

// third party payment type
var ThirdPaymentPayModelWechatPay = "WECHAT_PAY"

// third party payment service type
var ThirdPaymentServiceTypeProductOrder = "productOrder"

// TODO implement memberOrder
var ThirdPaymentServiceTypeMemberOrder = "memberOrder"
