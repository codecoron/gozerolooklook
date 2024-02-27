package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWonListCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWonListCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWonListCountLogic {
	return &GetWonListCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWonListCountLogic) GetWonListCount(in *pb.GetWonListCountReq) (*pb.GetWonListCountResp, error) {
	count, err := l.svcCtx.LotteryParticipationModel.GetWonListCountByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetWonListCountResp{
		Count: count,
	}, nil
}
