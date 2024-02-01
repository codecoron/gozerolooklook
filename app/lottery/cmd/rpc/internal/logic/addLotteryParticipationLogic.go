package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryParticipationLogic {
	return &AddLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLotteryParticipationLogic) AddLotteryParticipation(in *pb.AddLotteryParticipationReq) (*pb.AddLotteryParticipationResp, error) {
	if lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.LotteryId); err != nil {
		return nil, err
	} else if lottery.IsAnnounced != 0 {
		return nil, errors.New("抽奖已公布，不能参与")
	}

	r, err := l.svcCtx.LotteryParticipationModel.Insert(l.ctx, &model.LotteryParticipation{
		LotteryId: in.LotteryId,
		UserId:    in.UserId,
	})
	if err != nil {
		return nil, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &pb.AddLotteryParticipationResp{
		Id: id,
	}, nil
}
