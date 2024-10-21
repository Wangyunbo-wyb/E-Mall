package logic

import (
	"context"
	"database/sql"

	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"webshop/apps/product/rpc/internal/svc"
	"webshop/apps/product/rpc/product"
	"webshop/pkg/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockRevertLogic {
	return &DecrStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockRevertLogic) DecrStockRevert(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
	db, err := sqlx.NewMysql(l.svcCtx.Config.DataSource).RawDB()
	if err != nil {
		logx.Errorf("[DecrStockRevertLogic] failed to get db: %s", err)
		return nil, xerr.NewErrCode(xerr.DbError)
	}

	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		l.Logger.Errorf("[DecrStockRevertLogic] failed to get barrier: %s", err)
		return nil, xerr.NewErrCode(xerr.DtmError)
	}
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		_, err := l.svcCtx.ProductModel.TxUpdateStock(tx, in.Id, 1)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &product.DecrStockResponse{}, nil
}
