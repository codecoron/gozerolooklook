package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/app/checkin/cmd/rpc/internal/config"
	"looklook/app/checkin/model"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config              config.Config
	RedisClient         *redis.Redis
	CheckinRecordModel  model.CheckinRecordModel
	IntegralModel       model.IntegralModel
	IntegralRecordModel model.IntegralRecordModel
	TaskRecordModel     model.TaskRecordModel
	TasksModelModel     model.TasksModel

	UserCenterRpc usercenter.Usercenter
	CheckinRpc    checkin.Checkin
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		CheckinRecordModel:  model.NewCheckinRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		IntegralModel:       model.NewIntegralModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		IntegralRecordModel: model.NewIntegralRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TaskRecordModel:     model.NewTaskRecordModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		TasksModelModel:     model.NewTasksModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),

		CheckinRpc:    checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
		UserCenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
