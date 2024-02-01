package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetParticipationUserIdsByLotteryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetParticipationUserIdsByLotteryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParticipationUserIdsByLotteryIdLogic {
	return &GetParticipationUserIdsByLotteryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetParticipationUserIdsByLotteryIdLogic) GetParticipationUserIdsByLotteryId(in *pb.GetParticipationUserIdsByLotteryIdReq) (*pb.GetParticipationUserIdsByLotteryIdResp, error) {
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("lottery_id = ?", in.LotteryId)
	list, err := l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}

	resp := &pb.GetParticipationUserIdsByLotteryIdResp{}
	for i := range list {
		resp.UserIds = append(resp.UserIds, list[i].UserId)
	}

	return resp, nil
}
