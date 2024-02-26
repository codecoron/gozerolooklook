package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/comment/cmd/api/internal/config"
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	LotteryRpc    lottery.LotteryZrpcClient
	CommentRpc    comment.CommentZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		CommentRpc:    comment.NewCommentZrpcClient(zrpc.MustNewClient(c.CommentRpcConf)),
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UserCenterRpcConf)),
	}
}
