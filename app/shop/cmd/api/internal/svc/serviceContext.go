package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/shop/cmd/api/internal/config"
	"looklook/app/shop/model"
)

type ServiceContext struct {
	Config     config.Config
	GoodsModel model.GoodsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		GoodsModel: model.NewGoodsModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
