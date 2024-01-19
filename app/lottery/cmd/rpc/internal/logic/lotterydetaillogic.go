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

func (l *LotteryDetailLogic) LotteryDetail(in *pb.LotteryDetailReq) (*pb.LotteryDetailResp, error) {
	lotteryId := in.Id
	prizes, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, lotteryId)
	if err != nil {
		return nil, err
	}
	lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, lotteryId)
	if err != nil {
		return nil, err
	}
	var pbprizes []*pb.Prize
	for _, p := range prizes {
		var prize pb.Prize
		_ = copier.Copy(&prize, p)
		pbprizes = append(pbprizes, &prize)
	}
	resp := new(pb.LotteryDetailResp)
	resp.Lottery = new(pb.Lottery)
	err = copier.Copy(resp.Lottery, lottery)
	if err != nil {
		return nil, err
	}
	resp.Prizes = pbprizes
	return resp, nil

}
