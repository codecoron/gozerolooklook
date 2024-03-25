package logic

import (
	"context"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/common/constants"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryStatisticLogic {
	return &GetLotteryStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetLotteryStatistic 抽奖记录总数、发起抽奖记录总数、中奖记录总数RPC
func (l *GetLotteryStatisticLogic) GetLotteryStatistic(in *pb.GetLotteryStatisticReq) (*pb.GetLotteryStatisticResp, error) {
	// 从参与抽奖表中获取当前用户所有的抽奖记录
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("user_id = ?", in.UserId)
	ParticipationCount, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	// 从抽奖表获取当前用户发起的抽奖记录的总数
	builder = l.svcCtx.LotteryModel.SelectBuilder().Where("user_id = ?", in.UserId)
	CreatedCount, err := l.svcCtx.LotteryModel.FindCount(l.ctx, builder, "id")

	// 从参与抽奖表中获取当前用户所有的中奖记录总数
	builder = l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("user_id = ? and is_won = ?", in.UserId, constants.IsWon)
	WonCount, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")

	return &pb.GetLotteryStatisticResp{
		ParticipationCount: ParticipationCount,
		CreatedCount:       CreatedCount,
		WonCount:           WonCount,
	}, nil
}
