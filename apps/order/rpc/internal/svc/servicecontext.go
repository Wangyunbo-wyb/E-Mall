package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"webshop/apps/order/rpc/internal/config"
	"webshop/apps/order/rpc/internal/model"
	"webshop/apps/product/rpc/productclient"
	"webshop/apps/user/rpc/userclient"
)

type ServiceContext struct {
	Config         config.Config
	OrderModel     model.OrdersModel
	OrderitemModel model.OrderitemModel
	ShippingModel  model.ShippingModel
	UserRpc        userclient.User
	ProductRpc     productclient.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:         c,
		OrderModel:     model.NewOrdersModel(conn, c.CacheRedis),
		OrderitemModel: model.NewOrderitemModel(conn, c.CacheRedis),
		ShippingModel:  model.NewShippingModel(conn, c.CacheRedis),
		UserRpc:        userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc:     productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
