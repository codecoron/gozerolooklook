package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryLogic {
	return &SearchLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLotteryLogic) SearchLottery(in *pb.SearchLotteryReq) (*pb.SearchLotteryResp, error) {
	//list, err := l.svcCtx.LotteryModel.FindPageListByIdDESC(l.ctx, whereBuilder, in.LastId, in.PageSize)
	list, err := l.svcCtx.LotteryModel.LotteryList(l.ctx, in.Limit, in.IsSelected, in.LastId)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}

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
	return &pb.SearchLotteryResp{
		Lottery: resp,
	}, nil
}
