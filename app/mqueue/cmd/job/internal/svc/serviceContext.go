package svc

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/hibiken/asynq"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/mqueue/cmd/job/internal/config"
	"looklook/app/notice/cmd/rpc/notice"
	"looklook/app/order/cmd/rpc/order"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config        config.Config
	AsynqServer   *asynq.Server
	MiniProgram   *miniprogram.MiniProgram // looklook使用
	WxMiniProgram *miniProgram.MiniProgram // lottery使用
	CheckInRpc    checkin.Checkin
	OrderRpc      order.Order
	UsercenterRpc usercenter.Usercenter
	LotteryRpc    lottery.LotteryZrpcClient
	NoticeRpc     notice.Notice
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AsynqServer:   newAsynqServer(c),
		MiniProgram:   newMiniprogramClient(c), // looklook使用
		WxMiniProgram: MustNewMiniProgram(c),   // lottery使用

		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		LotteryRpc:    lottery.NewLotteryZrpcClient(zrpc.MustNewClient(c.LotteryRpcConf)),
		NoticeRpc:     notice.NewNotice(zrpc.MustNewClient(c.NoticeRpcConf)),
		CheckInRpc:    checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
	}
}
