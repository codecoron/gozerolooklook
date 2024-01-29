package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/mqueue/cmd/job/jobtype"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/wxnotice"
	"looklook/common/xerr"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrNotifyLotteryDrawFail = xerr.NewErrMsg("notify lottery draw fail")

type NoticeLotteryDrawLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeLotteryDrawLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeLotteryDrawLogic {
	return &NoticeLotteryDrawLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NoticeLotteryDrawLogic) NoticeLotteryDraw(in *pb.NoticeLotteryDrawReq) (*pb.NoticeLotteryDrawResp, error) {
	// 获取抽奖信息
	rpcLotteryInfo, err := l.svcCtx.LotteryRpc.GetLotteryById(l.ctx, &lottery.GetLotteryByIdReq{
		Id: in.LotteryId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to query the lottery"), "Failed to query the lottery, rpc GetLotteryById fail , lotteryId : %d , err : %v", in.LotteryId, err)
	}

	// 获取奖品信息
	rpcPrizeList, err := l.svcCtx.LotteryRpc.GetPrizeListByLotteryId(l.ctx, &lottery.GetPrizeListByLotteryIdReq{LotteryId: in.LotteryId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to query the prize"), "Failed to query the prize, rpc GetPrizeListByLotteryId fail , lotteryId : %d , err : %v", in.LotteryId, err)
	}

	// 汇总消息相关信息
	lotteryName := rpcLotteryInfo.Lottery.Name
	remindText := "看看你中奖了吗"

	prizeName, firstLevelPrizeName := "", ""
	if len(rpcPrizeList.Prizes) > 0 {
		for _, prize := range rpcPrizeList.Prizes {
			if prize.Level == 1 {
				firstLevelPrizeName = prize.Name
				break
			}
		}
		prizeName = firstLevelPrizeName + " 等"
	}

	for _, userId := range in.UserIds {
		userAuthResp, err := l.svcCtx.UserCenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
			UserId:   userId,
			AuthType: "wxMini",
		})
		if err != nil {
			return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeLotteryDrawLogic GetUserAuthByUserId err:%+v, userId:%d", err, userId)
		}
		if userAuthResp.UserAuth == nil || userAuthResp.UserAuth.AuthKey == "" {
			logx.WithContext(l.ctx).Errorw("NoticeLotteryDrawLogic user has no wechat auth",
				logx.Field("userId", userId))
			continue
		}
		openid := userAuthResp.UserAuth.AuthKey

		// 拼接小程序页面地址
		// TODO 地址需要规范化
		pageAddr := fmt.Sprintf("pages/detail/prize?lotterId=%d&userId=%d", rpcLotteryInfo.Lottery.Id, userId)

		msg := wxnotice.MessageLotteryDraw{
			PrizeName:   wxnotice.Item{Value: prizeName},
			LotteryName: wxnotice.Item{Value: lotteryName},
			RemindText:  wxnotice.Item{Value: remindText},
		}
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeLotteryDrawLogic msg json marshal err:%+v, msg:%+v", err, msg)
		}

		p := jobtype.WxMiniProgramNotifyUserPayload{
			MsgType:  msg.Type(),
			OpenId:   openid,
			PageAddr: pageAddr,
			Data:     string(jsonMsg),
		}

		payload, err := json.Marshal(p)
		if err != nil {
			return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeLotteryDrawLogic payload json marshal err:%+v, payload:%+v", err, p)
		}

		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgWxMiniProgramNotifyUser, payload))
		if err != nil {
			return nil, errors.Wrapf(ErrNotifyLotteryDrawFail, "NoticeLotteryDrawLogic AsynqClient.Enqueue err:%+v, payload:%+v", err, payload)
		}
	}

	return &pb.NoticeLotteryDrawResp{}, nil
}
