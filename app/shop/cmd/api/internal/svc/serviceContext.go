package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/shop/cmd/api/internal/config"
	"looklook/app/shop/cmd/rpc/shop"
	"looklook/app/shop/model"
)

type ServiceContext struct {
	Config     config.Config
	GoodsModel model.GoodsModel
	ShopRpc    shop.Shop
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		GoodsModel: model.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ShopRpc:    shop.NewShop(zrpc.MustNewClient(c.ShopRpcConf)),
	}
}
