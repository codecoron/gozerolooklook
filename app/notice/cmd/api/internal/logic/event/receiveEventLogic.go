package event

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"

	"looklook/app/notice/cmd/api/internal/svc"
	"looklook/app/notice/cmd/api/internal/types"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/zeromicro/go-zero/core/logx"
)

type ReceiveEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceiveEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReceiveEventLogic {
	return &ReceiveEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReceiveEventLogic) ReceiveEvent(req *types.ReceiveEventReq, r *http.Request) (resp *types.ReceiveEventResp, err error) {
	app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:  l.svcCtx.Config.WxMiniConf.AppId,
		Secret: l.svcCtx.Config.WxMiniConf.Secret,
		Token:  l.svcCtx.Config.WxMsgConf.EventToken,
		AESKey: l.svcCtx.Config.WxMsgConf.EncodingAESKey,
	})
	if err != nil {
		return nil, err
	}

	_, callbackMsgHeader, err := app.Server.GetEvent(r)
	if err != nil {
		return nil, err
	}

	// TODO 测试代码
	fmt.Println("content:", string(callbackMsgHeader.Content))

	var msg types.MsgEvent
	err = xml.Unmarshal(callbackMsgHeader.Content, &msg)
	if err != nil {
		return nil, err
	}

	// TODO 测试代码
	fmt.Printf("msg: %+v", msg)

	return
}
