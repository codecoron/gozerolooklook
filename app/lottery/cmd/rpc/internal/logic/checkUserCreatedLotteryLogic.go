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
	// todo 有两种错误，一种是没找到的错误，一种是其他错误，假如有其他错误返回其他错误，假如是没找到的错误，就返回0，正常返回1
	if id == nil {
		return &pb.CheckUserCreatedLotteryResp{IsCreated: 0}, nil
	}
	return &pb.CheckUserCreatedLotteryResp{IsCreated: 1}, nil
}
