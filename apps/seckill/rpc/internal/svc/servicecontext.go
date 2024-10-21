package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"webshop/apps/product/rpc/productclient"
	"webshop/apps/seckill/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	BizRedis    *redis.Redis
	ProductRPC  productclient.Product
	KafkaPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		BizRedis:    redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		ProductRPC:  productclient.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		KafkaPusher: kq.NewPusher(c.Kafka.Addrs, c.Kafka.SeckillTopic),
	}
}
