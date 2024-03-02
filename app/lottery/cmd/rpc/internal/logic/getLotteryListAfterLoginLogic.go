package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"

	"github.com/zeromicro/go-zero/core/logx"
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
	list, err := l.svcCtx.LotteryModel.LotteryList(l.ctx, in.Page, in.Size, in.IsSelected, in.LastId)
	if err != nil {
		return nil, err
	}

	// 获取当前用户参与过的抽奖
	ParticipatedLotteryIds, err := l.svcCtx.LotteryParticipationModel.GetParticipatedLotteryIdsByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	//fmt.Println("ParticipatedLotteryIds:", ParticipatedLotteryIds)

	// 当用户参与过之后，就将list中的这个抽奖去除
	var newList []*model.Lottery
	if len(ParticipatedLotteryIds) > 0 {
		// 将ParticipatedLotteryIds放入map中
		ParticipatedLotteryIdsMap := make(map[int64]bool)
		for _, lotteryId := range ParticipatedLotteryIds {
			ParticipatedLotteryIdsMap[lotteryId] = true
		}
		// 遍历list，将已经参与过的抽奖去除

		for _, lottery := range list {
			if _, ok := ParticipatedLotteryIdsMap[lottery.Id]; !ok {
				newList = append(newList, lottery)
			}
		}
	} else {
		newList = list

	}

	var resp []*pb.Lottery
	if len(newList) > 0 {
		for _, lottery := range newList {
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
