package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	BizRedis   redis.RedisConf
	CacheRedis cache.CacheConf
	OrderRPC   zrpc.RpcClientConf
	Mysql      struct {
		DataSource string
	}
	Kafka struct {
		Addrs        []string
		PaymentTopic string
	}
}
