package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
)

type GetLotteryListAfterLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryListAfterLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListAfterLoginLogic {
	return &GetLotteryListAfterLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryListAfterLoginLogic) GetLotteryListAfterLogin(in *pb.GetLotteryListAfterLoginReq) (*pb.GetLotteryListAfterLoginResp, error) {
	// 获取当前用户参与过的抽奖
	ParticipatedLotteryIds, err := l.svcCtx.LotteryParticipationModel.GetParticipatedLotteryIdsByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	if in.LastId == 0 {
		in.LastId, err = l.svcCtx.LotteryModel.GetLastId(l.ctx)
		if err != nil {
			return nil, err
		}
	}

	list, err := l.svcCtx.LotteryModel.GetLotteryListAfterLogin(l.ctx, in.Size, in.IsSelected, in.LastId, ParticipatedLotteryIds)
	if err != nil {
		return nil, err
	}

	// 当用户参与过之后，就将list中的这个抽奖去除

	var resp []*pb.Lottery
	if len(list) > 0 {
		for _, lottery := range list {
			var pbLottery pb.Lottery
			_ = copier.Copy(&pbLottery, lottery)
			pbLottery.PublishTime = lottery.PublishTime.Time.Unix()
			pbLottery.AwardDeadline = lottery.AwardDeadline.Unix()
			pbLottery.AnnounceType = lottery.AnnounceType
			pbLottery.AnnounceTime = lottery.AnnounceTime.Unix()
			pbLottery.IsAnnounced = lottery.IsAnnounced
			resp = append(resp, &pbLottery)
		}
	}

	return &pb.GetLotteryListAfterLoginResp{
		List: resp,
	}, nil
}
