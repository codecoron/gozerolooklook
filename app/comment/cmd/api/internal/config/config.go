package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	CommentRpcConf zrpc.RpcClientConf
	DB             struct {
		DataSource string
	}
	JwtAuth struct {
		AccessSecret string
	}
	Cache cache.CacheConf

	UserCenterRpcConf zrpc.RpcClientConf
	LotteryRpcConf    zrpc.RpcClientConf
}
