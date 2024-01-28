package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByLotteryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByLotteryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByLotteryIdLogic {
	return &FindByLotteryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByLotteryIdLogic) FindByLotteryId(in *pb.FindByLotteryIdReq) (*pb.FindByLotteryIdResp, error) {
	res, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, err
	}
	prizes := make([]*pb.Prize, 0)
	for _, p := range res {
		pbPrize := new(pb.Prize)
		err := copier.Copy(pbPrize, p)
		if err != nil {
			return nil, err
		}
		prizes = append(prizes, pbPrize)
	}
	return &pb.FindByLotteryIdResp{
		Prizes: prizes,
	}, nil
}
