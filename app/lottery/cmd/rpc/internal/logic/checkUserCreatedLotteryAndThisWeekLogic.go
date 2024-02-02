package logic

import (
	"context"
	"fmt"
	"looklook/common/constants"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckUserCreatedLotteryAndThisWeekLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCreatedLotteryAndThisWeekLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCreatedLotteryAndThisWeekLogic {
	return &CheckUserCreatedLotteryAndThisWeekLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCreatedLotteryAndThisWeekLogic) CheckUserCreatedLotteryAndThisWeek(in *pb.CheckUserCreatedLotteryAndThisWeekReq) (*pb.CheckUserCreatedLotteryAndThisWeekResp, error) {
	userId := in.UserId
	// 根据uid获取当前用户并且今天发布的的所有抽奖id
	LotteryIds, err := l.svcCtx.LotteryModel.GetWeekLotteryIdsByUserId(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	fmt.Println("lotterys:", LotteryIds)
	// 判断是否有一个抽奖符合，有一个符合就跳出循环，返回yes = 1
	for _, lotteryId := range LotteryIds {
		yes, err := l.CheckLotteryThisWeek(lotteryId)
		if err != nil {
			return nil, err
		}
		if yes {
			return &pb.CheckUserCreatedLotteryAndThisWeekResp{
				Yes: 1,
			}, nil
		}
	}
	return &pb.CheckUserCreatedLotteryAndThisWeekResp{
		Yes: 0,
	}, nil
}

// CheckLotteryToday 检查抽奖是否在今天之内发起并有超过五个人参加
func (l *CheckUserCreatedLotteryAndThisWeekLogic) CheckLotteryThisWeek(lotteryID int64) (bool, error) {
	participantsCount, err := l.svcCtx.LotteryParticipationModel.GetParticipatorsCountByLotteryId(l.ctx, lotteryID)
	if err != nil {
		return false, err
	}
	fmt.Println("participantsCount:", participantsCount)
	// 判断抽奖是否在今天之内发起并有超过五个人参加
	return participantsCount > constants.LotteryThisWeekParticipantsCount, nil
}
