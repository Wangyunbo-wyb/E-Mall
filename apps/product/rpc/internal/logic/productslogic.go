package logic

import (
	"context"
	"strconv"
	"strings"

	"webshop/apps/product/rpc/internal/model"
	"webshop/apps/product/rpc/internal/svc"
	"webshop/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type ProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductsLogic {
	return &ProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductsLogic) Products(in *product.ProductRequest) (*product.ProductResponse, error) {
	products := make(map[int64]*product.ProductItem)
	pdis := strings.Split(in.ProductIds, ",")
	ps, err := mr.MapReduce(func(source chan<- string) {
		for _, pid := range pdis {
			source <- pid
		}
	}, func(item string, writer mr.Writer[*model.Product], cancel func(error)) {
		pid, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			return
		}
		p, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(pid))
		if err != nil {
			return
		}
		writer.Write(p)
	}, func(pipe <-chan *model.Product, writer mr.Writer[[]*model.Product], cancel func(error)) {
		var r []*model.Product
		for p := range pipe {
			r = append(r, p)
		}
		writer.Write(r)
	})
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		products[int64(p.Id)] = &product.ProductItem{
			ProductId: int64(p.Id),
			Name:      p.Name,
		}
	}
	return &product.ProductResponse{Products: products}, nil
}
