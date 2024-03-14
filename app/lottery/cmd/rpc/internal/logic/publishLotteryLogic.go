package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLotteryLogic {
	return &PublishLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLotteryLogic) PublishLottery(in *pb.PublishLotteryReq) (*pb.PublishLotteryResp, error) {
	// todo: add your logic here and delete this line

	return &pb.PublishLotteryResp{}, nil
}
