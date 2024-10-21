package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"webshop/apps/pay/rpc/internal/config"
	"webshop/apps/pay/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	ThirdPaymentModel model.ThirdPaymentModel
	KafkaPusher       *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:            c,
		ThirdPaymentModel: model.NewThirdPaymentModel(sqlConn, c.CacheRedis),
		KafkaPusher:       kq.NewPusher(c.Kafka.Addrs, c.Kafka.PaymentTopic),
	}
}
