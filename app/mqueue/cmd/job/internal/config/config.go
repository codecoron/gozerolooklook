package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	Redis      redis.RedisConf
	WxMiniConf WxMiniConf
	WxMsgConf  WxMsgConf

	SettleRpcConf     zrpc.RpcClientConf
	OrderRpcConf      zrpc.RpcClientConf
	UsercenterRpcConf zrpc.RpcClientConf
	LotteryRpcConf    zrpc.RpcClientConf
}
