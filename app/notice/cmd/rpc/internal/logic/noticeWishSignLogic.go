package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/wxnotice"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeWishSignLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeWishSignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeWishSignLogic {
	return &NoticeWishSignLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NoticeWishSignLogic) NoticeWishSign(in *pb.NoticeWishSignInReq) (*pb.NoticeWishSignInResp, error) {
	//获取用户信息
	userAuthResp, err := l.svcCtx.UserCenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
		UserId:   in.UserId,
		AuthType: "wxMini",
	})
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeWishSign GetUserAuthByUserId err:%+v, userId:%d", err, in.UserId)
	}
	if userAuthResp.UserAuth == nil || userAuthResp.UserAuth.AuthKey == "" {
		logx.WithContext(l.ctx).Errorw("NoticeWishSign user has no wechat auth",
			logx.Field("userId", in.UserId))
	}
	openid := userAuthResp.UserAuth.AuthKey

	// 拼接小程序页面地址
	pageAddr := fmt.Sprintf("pages/detail/prize?lotterId=%d&userId=%d", in.UserId, in.UserId)

	//构建微信消息
	checkinType := "每日心愿签到"
	remindText := "别忘记签到哦"
	msg := wxnotice.MessageWishCheckin{
		CheckinType: wxnotice.Item{Value: checkinType},
		Reward:      wxnotice.Item{Value: wxnotice.ConvertToChineseNumber(in.Reward)},
		Accumulate:  wxnotice.Item{Value: wxnotice.ConvertToChineseNumber(in.Accumulate)},
		RemindText:  wxnotice.Item{Value: remindText},
	}
	marshal, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeWishSign msg json marshal err:%+v, msg:%+v", err, msg)
	}
	//封装asynq消息
	p := jobtype.WxMiniProgramNotifyUserPayload{
		MsgType:  msg.Type(),
		OpenId:   openid,
		PageAddr: pageAddr,
		Data:     string(marshal),
	}
	payload, err := json.Marshal(p)
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeWishSign payload json marshal err:%+v, payload:%+v", err, p)
	}
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgWxMiniProgramNotifyUser, payload))
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeWishSign AsynqClient.Enqueue err:%+v, payload:%+v", err, payload)
	}
	return &pb.NoticeWishSignInResp{}, nil
}
