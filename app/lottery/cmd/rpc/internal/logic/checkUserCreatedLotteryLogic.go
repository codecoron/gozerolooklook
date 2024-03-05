package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
)

type CheckUserCreatedLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserCreatedLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserCreatedLotteryLogic {
	return &CheckUserCreatedLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckUserCreatedLotteryLogic) CheckUserCreatedLottery(in *pb.CheckUserCreatedLotteryReq) (*pb.CheckUserCreatedLotteryResp, error) {
	id, err := l.svcCtx.LotteryModel.GetLotteryIdByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	if id == nil {
		return &pb.CheckUserCreatedLotteryResp{IsCreated: 0}, nil
	}
	return &pb.CheckUserCreatedLotteryResp{IsCreated: 1}, nil
}
