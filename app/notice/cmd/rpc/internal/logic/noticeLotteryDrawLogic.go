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
	// TODO 还需要对接，改用俊威的GetPrizeListByLotteryId
	rpcPrizeList, err := l.svcCtx.LotteryRpc.SearchPrize(l.ctx, &lottery.SearchPrizeReq{LotteryId: in.LotteryId})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to query the prize"), "Failed to query the prize, rpc SearchPrize fail , lotteryId : %d , err : %v", in.LotteryId, err)
	}
	if rpcPrizeList == nil {
		return &pb.NoticeLotteryDrawResp{}, nil
	}

	// 汇总消息相关信息
	lotteryName := rpcLotteryInfo.Lottery.Name
	remindText := "看看你中奖了吗"
	firstLevelPrizeName := ""
	for _, prize := range rpcPrizeList.Prize {
		if prize.Level == 1 {
			firstLevelPrizeName = prize.Name
			break
		}
	}

	for _, userId := range in.UserIds {
		userAuthResp, err := l.svcCtx.UserCenterRpc.GetUserAuthByUserId(l.ctx, &usercenter.GetUserAuthByUserIdReq{
			UserId:   userId,
			AuthType: "wxMini",
		})
		if err != nil {
			logx.WithContext(l.ctx).Error("NoticeLotteryDrawLogic UserCenterRpc.GetUserAuthByUserId err",
				logx.Field("err", err))
			continue
		}
		if userAuthResp.UserAuth == nil || userAuthResp.UserAuth.AuthKey == "" {
			logx.WithContext(l.ctx).Error("NoticeLotteryDrawLogic user has no wechat auth",
				logx.Field("userId", userId))
			continue
		}

		// 拼接小程序页面地址
		// TODO 地址需要规范化
		pageAddr := fmt.Sprintf("pages/home/index?lottery_id=%d&userId=%d", rpcLotteryInfo.Lottery.Id, userId)

		openid := userAuthResp.UserAuth.AuthKey
		// TODO 测试代码
		//prizeName := firstLevelPrizeName + " 等"
		prizeName := "一等奖XXX"
		fmt.Println(firstLevelPrizeName)

		msg := wxnotice.MessageLotteryDraw{
			PrizeName:   wxnotice.Item{Value: prizeName},
			LotteryName: wxnotice.Item{Value: lotteryName},
			RemindText:  wxnotice.Item{Value: remindText},
		}
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			logx.WithContext(l.ctx).Error("NoticeLotteryDrawLogic msg json marshal err",
				logx.Field("err", err),
				logx.Field("msg", msg))
			continue
		}

		p := jobtype.WxMiniProgramNotifyUserPayload{
			MsgType:  msg.Type(),
			OpenId:   openid,
			PageAddr: pageAddr,
			Data:     string(jsonMsg),
		}

		payload, err := json.Marshal(p)
		if err != nil {
			logx.WithContext(l.ctx).Error("NoticeLotteryDrawLogic payload json marshal err",
				logx.Field("err", err),
				logx.Field("payload", p))
			continue
		}

		_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgWxMiniProgramNotifyUser, payload))
		if err != nil {
			logx.WithContext(l.ctx).Error("NoticeLotteryDrawLogic AsynqClient.Enqueue err",
				logx.Field("err", err),
				logx.Field("payload", payload))
			continue
		}
	}

	return &pb.NoticeLotteryDrawResp{}, nil
}
