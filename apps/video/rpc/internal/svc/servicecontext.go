package svc

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"webshop/apps/video/rpc/internal/config"
	"webshop/pkg/orm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *orm.DB
	Redis       *redis.Client
	AsynqClient *asynq.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.NewMysql(&orm.Config{
		DSN:         c.DB.DataSource,
		Active:      100,
		Idle:        20,
		IdleTimeout: time.Hour,
	})

	rds := redis.NewClient(&redis.Options{
		Addr:     c.BizRedis.Host,
		Password: c.BizRedis.Pass,
	})
	return &ServiceContext{
		Config:      c,
		DB:          db,
		Redis:       rds,
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{Addr: c.BizRedis.Host, Password: c.BizRedis.Pass}),
	}
}
