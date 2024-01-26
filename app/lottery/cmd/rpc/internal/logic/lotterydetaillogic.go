package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LotteryDetailLogic) LotteryDetail(in *pb.LotteryDetailReq) (resp *pb.LotteryDetailResp, err error) {
	lotteryId := in.Id
	res, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, lotteryId)
	if err != nil {
		return nil, err
	}
	lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, lotteryId)
	if err != nil {
		return nil, err
	}
	resp = new(pb.LotteryDetailResp)
	resp.Lottery = new(pb.Lottery)
	_ = copier.Copy(resp.Lottery, lottery)
	for _, p := range res {
		prize := new(pb.Prize)
		_ = copier.Copy(prize, p)
		resp.Prizes = append(resp.Prizes, prize)
	}
	return resp, nil

}
