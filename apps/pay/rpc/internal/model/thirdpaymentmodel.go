package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ThirdPaymentModel = (*customThirdPaymentModel)(nil)

var (
	cachePaymentThirdPaymentIdPrefix = "cache#Payment#thirdPayment#id#"
	cachePaymentThirdPaymentSnPrefix = "cache#Payment#thirdPayment#sn#"
)

type (
	// ThirdPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThirdPaymentModel.
	ThirdPaymentModel interface {
		thirdPaymentModel
		UpdateWithVersion(ctx context.Context, session sqlx.Session, data *ThirdPayment) error
	}

	customThirdPaymentModel struct {
		*defaultThirdPaymentModel
	}
)

// NewThirdPaymentModel returns a model for the database table.
func NewThirdPaymentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ThirdPaymentModel {
	return &customThirdPaymentModel{
		defaultThirdPaymentModel: newThirdPaymentModel(conn, c, opts...),
	}
}

func (m *defaultThirdPaymentModel) UpdateWithVersion(ctx context.Context, session sqlx.Session, newData *ThirdPayment) error {

	oldVersion := newData.Version
	newData.Version += 1

	var sqlResult sql.Result
	var err error

	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}
	PaymentThirdPaymentIdKey := fmt.Sprintf("%s%v", cachePaymentThirdPaymentIdPrefix, data.Id)
	PaymentThirdPaymentSnKey := fmt.Sprintf("%s%v", cachePaymentThirdPaymentSnPrefix, data.Sn)
	sqlResult, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? and version = ? ", m.table, thirdPaymentRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, newData.Sn, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.PayMode, newData.TradeType, newData.TradeState, newData.PayTotal, newData.TransactionId, newData.TradeStateDesc, newData.OrderId, newData.ServiceType, newData.PayStatus, newData.PayTime, newData.Id, oldVersion)
		}
		return conn.ExecCtx(ctx, query, newData.Sn, newData.DeleteTime, newData.DelState, newData.Version, newData.UserId, newData.PayMode, newData.TradeType, newData.TradeState, newData.PayTotal, newData.TransactionId, newData.TradeStateDesc, newData.OrderId, newData.ServiceType, newData.PayStatus, newData.PayTime, newData.Id, oldVersion)
	}, PaymentThirdPaymentIdKey, PaymentThirdPaymentSnKey)
	if err != nil {
		return err
	}
	updateCount, err := sqlResult.RowsAffected()
	if err != nil {
		return err
	}
	if updateCount == 0 {
		return ErrNoRowsUpdate
	}

	return nil
}
