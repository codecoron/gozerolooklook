package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/lottery/cmd/rpc/internal/config"
	"looklook/app/lottery/model"
	"looklook/app/notice/cmd/rpc/notice"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config                    config.Config
	RedisClient               *redis.Redis
	LotteryModel              model.LotteryModel
	PrizeModel                model.PrizeModel
	ClockTaskModel            model.ClockTaskModel
	ClockTaskRecordModel      model.ClockTaskRecordModel
	UserCenterRpc             usercenter.Usercenter
	NoticeRpc                 notice.Notice
	LotteryParticipationModel model.LotteryParticipationModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                    c,
		LotteryModel:              model.NewLotteryModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		PrizeModel:                model.NewPrizeModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		UserCenterRpc:             usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
		// NoticeRpc:                 notice.NewNotice(zrpc.MustNewClient(c.NoticeRpcConf)),
		LotteryParticipationModel: model.NewLotteryParticipationModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ClockTaskModel:            model.NewClockTaskModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		ClockTaskRecordModel:      model.NewClockTaskRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
