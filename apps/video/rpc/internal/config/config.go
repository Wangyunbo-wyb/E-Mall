package config

import (
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DataSource  string
		Active      int           `json:",default=10"`
		Idle        int           `json:",default=100"`
		IdleTimeout time.Duration `json:",default=3600s"`
	}
	BizRedis   redis.RedisConf
	CacheRedis cache.CacheConf
}
