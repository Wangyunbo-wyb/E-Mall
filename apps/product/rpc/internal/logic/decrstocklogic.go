package logic

import (
	"context"
	"database/sql"
	"errors"

	"github.com/dtm-labs/dtmcli"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"webshop/apps/product/rpc/internal/svc"
	"webshop/apps/product/rpc/product"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockLogic {
	return &DecrStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockLogic) DecrStock(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.DataSource).RawDB()
	if err != nil {
		logx.Errorf("[DecrStockLogic] failed to get db: %s", err)
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	// get barrier from dtm
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		l.Logger.Errorf("[DecrStockLogic] failed to get barrier: %s", err)
		return nil, xerr.NewErrCode(xerr.DtmError)
	}
	// start a sub-transaction
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// update the stock
		result, err := l.svcCtx.ProductModel.TxUpdateStock(tx, in.Id, -1)
		if err != nil {
			return err
		}

		affected, err := result.RowsAffected()
		// if no rows affected, it means the stock is not enough
		if err == nil && affected == 0 {
			return dtmcli.ErrFailure
		}

		return err
	})

	// the transaction is aborted
	if errors.Is(err, dtmcli.ErrFailure) {
		l.Logger.Errorf("[DecrStockLogic] dtm transaction failed: %s", err)
		return nil, xerr.NewErrCode(xerr.DtmError)
	}

	if err != nil {
		return nil, err
	}

	return &product.DecrStockResponse{}, nil
}
