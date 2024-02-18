package logic

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/json"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/wxnotice"
	"looklook/common/xerr"
	"time"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeLotteryStartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeLotteryStartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeLotteryStartLogic {
	return &NoticeLotteryStartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NoticeLotteryStartLogic) NoticeLotteryStart(in *pb.NoticeLotteryStartReq) (*pb.NoticeLotteryStartResp, error) {
	// 获取抽奖信息
	rpcLotteryInfo, err := l.svcCtx.LotteryRpc.GetLotteryById(l.ctx, &lottery.GetLotteryByIdReq{
		Id: in.LotteryId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to query the lottery"), "Failed to query the lottery, rpc GetLotteryById fail , lotteryId : %d , err : %v", in.LotteryId, err)
	}
	//获取用户信息
	userAuthResp, err := l.svcCtx.UserCenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
		UserId:   in.UserId,
		AuthType: "wxMini",
	})
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeLotteryStart GetUserAuthByUserId err:%+v, userId:%d", err, in.UserId)
	}
	if userAuthResp.UserAuth == nil || userAuthResp.UserAuth.AuthKey == "" {
		logx.WithContext(l.ctx).Errorw("NoticeLotteryStart user has no wechat auth",
			logx.Field("userId", in.UserId))
	}
	openid := userAuthResp.UserAuth.AuthKey
	// 拼接小程序页面地址
	pageAddr := fmt.Sprintf("pages/detail/prize?lotterId=%d&userId=%d", rpcLotteryInfo.Lottery.Id, in.UserId)

	//构建微信消息
	remindText := "抽奖即将开始"
	msg := wxnotice.MessageLotteryStart{
		LotteryName: wxnotice.Item{Value: rpcLotteryInfo.GetLottery().Name},
		RemindText:  wxnotice.Item{Value: remindText},
	}
	marshal, err := json.Marshal(msg)
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeLotteryStart msg json marshal err:%+v, msg:%+v", err, msg)
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
		return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeLotteryStart payload json marshal err:%+v, payload:%+v", err, p)
	}

	//计算发送时间
	startTime := time.Unix(in.StartTime, 0)
	nowTime := time.Now()
	sub := startTime.Sub(nowTime)
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgWxMiniProgramNotifyUser, payload), asynq.ProcessIn(time.Duration(sub)))
	if err != nil {
		return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeLotteryStart AsynqClient.Enqueue err:%+v, payload:%+v", err, payload)
	}
	//发送消息
	return &pb.NoticeLotteryStartResp{}, nil
}
