package logic

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"looklook/common/constants"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckSelectedLotteryParticipatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckSelectedLotteryParticipatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckSelectedLotteryParticipatedLogic {
	return &CheckSelectedLotteryParticipatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckSelectedLotteryParticipatedLogic) CheckSelectedLotteryParticipated(in *pb.CheckSelectedLotteryParticipatedReq) (*pb.CheckSelectedLotteryParticipatedResp, error) {
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().
		Where("user_id = ?", in.UserId)
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
	count, err := l.svcCtx.LotteryParticipationModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}
	if count > 0 {
		count = 1
	}
	return &pb.CheckSelectedLotteryParticipatedResp{
		Participated: count,
	}, nil
}
