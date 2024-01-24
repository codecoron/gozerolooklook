package svc

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"looklook/app/notice/cmd/api/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	WxMiniProgram *miniprogram.MiniProgram
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		WxMiniProgram: wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
			AppID:     c.WxMiniConf.AppId,
			AppSecret: c.WxMiniConf.Secret,
			Cache:     cache.NewMemory(),
		}),
	}
}
