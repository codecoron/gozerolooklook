package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"
	"time"
)

type AddLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryLogic {
	return &AddLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------发起抽奖----------------------
func (l *AddLotteryLogic) AddLottery(in *pb.AddLotteryReq) (*pb.AddLotteryResp, error) {
	//TODO 添加事务处理
	//抽奖基本信息
	lottery := new(model.Lottery)
	lottery.UserId = in.UserId
	lottery.Name = in.Name
	lottery.AwardDeadline = time.Unix(in.AwardDeadline, 0)
	lottery.Introduce = in.Introduce
	lottery.JoinNumber = in.JoinNumber
	lottery.PublishType = in.PublishType
	lottery.Thumb = in.Thumb
	insert, err := l.svcCtx.LotteryModel.Insert(l.ctx, lottery)
	if err != nil {
		return nil, err
	}
	lotteryId, _ := insert.LastInsertId()
	//添加奖品信息
	for _, pbPrize := range in.Prizes {
		prize := new(model.Prize)
		err := copier.Copy(&prize, pbPrize)
		if err != nil {
			return nil, err
		}
		prize.LotteryId = lotteryId
		_, err = l.svcCtx.PrizeModel.Insert(l.ctx, prize)
		if err != nil {
			return nil, err
		}
	}
	return &pb.AddLotteryResp{
		Id: lotteryId,
	}, nil
}
