package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/vote/cmd/api/internal/config"
	"looklook/app/vote/cmd/rpc/vote"
)

type ServiceContext struct {
	Config  config.Config
	VoteRpc vote.Vote
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		VoteRpc: vote.NewVote(zrpc.MustNewClient(c.VoteRpcConf)),
	}
}
