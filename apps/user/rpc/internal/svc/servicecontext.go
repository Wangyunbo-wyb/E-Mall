package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"webshop/apps/user/rpc/internal/config"
	"webshop/apps/user/rpc/internal/model"
)

type ServiceContext struct {
	Config                  config.Config
	UserModel               model.UserModel
	UserReceiveAddressModel model.UserReceiveAddressModel
	UserCollectionModel     model.UserCollectionModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                  c,
		UserModel:               model.NewUserModel(sqlConn, c.CacheRedis),
		UserReceiveAddressModel: model.NewUserReceiveAddressModel(sqlConn, c.CacheRedis),
		UserCollectionModel:     model.NewUserCollectionModel(sqlConn, c.CacheRedis),
	}
}
