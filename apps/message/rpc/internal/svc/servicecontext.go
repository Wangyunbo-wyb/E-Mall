package svc

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
	"webshop/apps/message/rpc/internal/config"
	"webshop/pkg/orm"
)

type ServiceContext struct {
	Config      config.Config
	DB          *orm.DB
	Redis       *redis.Client
	AsynqClient *asynq.Client
}

type DBList struct {
	Mysql *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := orm.NewMysql(&orm.Config{
		DSN:         c.DB.DataSource,
		Active:      c.DB.MaxOpenConns,
		Idle:        c.DB.MaxIdleConns,
		IdleTimeout: time.Duration(c.DB.MaxLifetime),
	})

	rds := redis.NewClient(&redis.Options{
		Addr:         c.BizRedis.Address,
		Password:     c.BizRedis.Password,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		PoolTimeout:  3 * time.Second,
	})

	return &ServiceContext{
		Config:      c,
		DB:          db,
		Redis:       rds,
		AsynqClient: asynq.NewClient(asynq.RedisClientOpt{Addr: c.BizRedis.Address, Password: c.BizRedis.Password}),
	}
}
