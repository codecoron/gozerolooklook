package svc

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/notice/cmd/api/internal/config"
	"looklook/app/notice/cmd/rpc/notice"
)

type ServiceContext struct {
	Config        config.Config
	WxMiniProgram *miniProgram.MiniProgram

	NoticeRpc notice.Notice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		WxMiniProgram: MustNewMiniProgram(c),

		NoticeRpc: notice.NewNotice(zrpc.MustNewClient(c.NoticeRpcConf)),
	}
}
