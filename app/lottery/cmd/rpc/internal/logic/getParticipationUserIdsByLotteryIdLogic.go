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
	// todo: add your logic here and delete this line

	return &pb.GetParticipationUserIdsByLotteryIdResp{}, nil
}
