package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/lottery/cmd/rpc/internal/config"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/lottery/model"
)

type ServiceContext struct {
	Config       config.Config
	LotteryRpc   lottery.LotteryZrpcClient
	LotteryModel model.LotteryModel
	PrizeModel   model.PrizeModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		LotteryRpc:   lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		LotteryModel: model.NewLotteryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		PrizeModel:   model.NewPrizeModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
