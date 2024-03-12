package svc

import (
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/api/internal/config"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config                config.Config
	UsercenterRpc         usercenter.Usercenter
	CommentRpcConf        comment.CommentZrpcClient
	LotteryRpcConf        lottery.LotteryZrpcClient
	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:         c,
		UsercenterRpc:  usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		CommentRpcConf: comment.NewCommentZrpcClient(zrpc.MustNewClient(c.CommentRpcConf)),
		LotteryRpcConf: lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
	}
}
