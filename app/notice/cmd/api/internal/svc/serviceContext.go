package svc

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"looklook/app/notice/cmd/api/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	WxMiniProgram *miniProgram.MiniProgram
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		WxMiniProgram: MustNewMiniProgram(c),
	}
}
