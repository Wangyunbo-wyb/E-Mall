package service

import (
	"sync"

	"webshop/apps/order/rpc/orderclient"
	"webshop/apps/product/rpc/productclient"
	"webshop/apps/seckill/rmq/internal/config"
)

const (
	chanCount   = 10
	bufferCount = 1024
)

type Service struct {
	c          config.Config
	ProductRPC productclient.Product
	OrderRPC   orderclient.Order

	waiter   sync.WaitGroup
	msgsChan []chan *KafkaData
}
