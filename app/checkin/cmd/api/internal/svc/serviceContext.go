package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/checkin/cmd/api/internal/config"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	CheckinRpc    checkin.Checkin
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
		CheckinRpc:    checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
	}
}
