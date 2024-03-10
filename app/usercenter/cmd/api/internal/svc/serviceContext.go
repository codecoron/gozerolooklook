package svc

import (
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/api/internal/config"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	UsercenterRpc         usercenter.Usercenter
	LotteryRpc            lottery.LotteryZrpcClient
	CheckinRpc            checkin.Checkin
	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		CheckinRpc:    checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
	}
}
