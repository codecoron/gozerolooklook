package logic

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/common/constants"

	"github.com/jinzhu/now"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSelectedLotteryStatisticLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSelectedLotteryStatisticLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSelectedLotteryStatisticLogic {
	return &GetSelectedLotteryStatisticLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSelectedLotteryStatisticLogic) GetSelectedLotteryStatistic(in *pb.GetSelectedLotteryStatisticReq) (*pb.GetSelectedLotteryStatisticResp, error) {
	start := now.BeginningOfDay()
	end := now.EndOfDay()
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().
		Where("user_id = ? AND create_time >= ? AND create_time <= ?", in.UserId, start, end)
	participations, err := l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}
	lotteryIds := make([]int64, len(participations))
	for i := range participations {
		lotteryIds[i] = participations[i].LotteryId
	}
	builder = l.svcCtx.LotteryModel.SelectBuilder().
		Where(sq.Eq{"id": lotteryIds}).
		Where("is_selected = ?", constants.IsSelectedLottery)
	dayCount, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	start = now.BeginningOfWeek()
	end = now.EndOfWeek()
	builder = l.svcCtx.LotteryParticipationModel.SelectBuilder().
		Where("user_id = ? AND create_time >= ? AND create_time <= ?", in.UserId, start, end)
	participations, err = l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}
	lotteryIds = make([]int64, len(participations))
	for i := range participations {
		lotteryIds[i] = participations[i].LotteryId
	}
	builder = l.svcCtx.LotteryModel.SelectBuilder().
		Where(sq.Eq{"id": lotteryIds}).
		Where("is_selected = ?", constants.IsSelectedLottery)
	weekCount, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}

	resp := &pb.GetSelectedLotteryStatisticResp{
		DayCount:  dayCount,
		WeekCount: weekCount,
	}

	return resp, nil
}
