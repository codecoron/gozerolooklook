package logic

import (
	"context"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	model2 "looklook/app/lottery/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryPrizesListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryPrizesListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryPrizesListByUserIdLogic {
	return &GetLotteryPrizesListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

/**
 * message LotteryPrizes {
  int64 lotteryId = 2; // 抽奖id
  repeated Prize Prizes = 3; // 奖品列表
  int64 ParticipateTime = 4; // 参与时间
  int64 IsWon = 5; // 是否中奖
  int64 CreateTime = 6; // 创建时间
}
*/

type LotteryPrizes struct {
	lotteryId       int64
	Prizes          []*model2.Prize
	ParticipateTime int64
	IsWon           int64
	CreateTime      int64
	WonTime         int64
}

func (l *GetLotteryPrizesListByUserIdLogic) GetLotteryPrizesListByUserId(in *pb.GetLotteryPrizesListByUserIdReq) (*pb.GetLotteryPrizesListByUserIdResp, error) {
	var err error
	lotterys := make([]*model2.Lottery3, 0)

	switch in.Type {
	case 1:
		if in.LastId == 0 {
			lastId, err := l.svcCtx.LotteryModel.GetLastId(l.ctx)
			if err != nil {
				return nil, err
			}
			in.LastId = lastId + 1
		}
		// 获取当前用户所有参与的抽奖信息
		//fmt.Println("获取当前用户所有参与的抽奖信息", in.UserId, in.LastId, in.Size, in.IsAnnounced, in.Type)
		lotterys2, err := l.svcCtx.LotteryParticipationModel.FindAllByUserId(in.UserId, in.LastId, in.Size, in.IsAnnounced)
		if err != nil {
			return nil, err
		}
		for _, lottery := range lotterys2 {
			lotterys = append(lotterys, &model2.Lottery3{
				Id:   lottery.Id,
				Time: lottery.Time,
			})
		}

	case 2:
		// 获取当前用户发起的所有抽奖
		if in.LastId == 0 {
			lastId, err := l.svcCtx.LotteryModel.GetLastId(l.ctx)
			if err != nil {
				return nil, err
			}
			in.LastId = lastId + 1
		}
		//fmt.Println("获取当前用户发起的所有抽奖", in.UserId, in.LastId, in.Size, in.IsAnnounced, in.Type)
		lotterys2, err := l.svcCtx.LotteryModel.FindAllByUserId(in.UserId, in.LastId, in.Size, in.IsAnnounced)
		if err != nil {
			return nil, err
		}
		for _, lottery := range lotterys2 {
			lotterys = append(lotterys, &model2.Lottery3{
				Id:   lottery.Id,
				Time: lottery.Time,
			})
		}
	case 3:

		if in.LastId == 0 {
			lastId, err := l.svcCtx.LotteryParticipationModel.GetLastId(l.ctx)
			if err != nil {
				return nil, err
			}
			in.LastId = lastId + 1
		}
		// 获取当前用户所有参与的抽奖信息（中奖的）
		//fmt.Println("获取当前用户所有中奖的抽奖信息", in.UserId, in.LastId, in.Size, in.IsAnnounced, in.Type)
		// 分页根据参与表的id作为lastId
		lotterys, err = l.svcCtx.LotteryParticipationModel.FindWonListByUserId(in.UserId, in.LastId, in.Size, in.IsAnnounced)
		if err != nil {
			return nil, err
		}
	}

	//for _, lottery := range lotterys {
	//	fmt.Println("获取当前用户抽奖列表：", in.Type, lottery)
	//}

	// 得到LottoryList后，根据lotteryId获取奖品列表
	// 先统计lotteryIds
	lotteryIds := make([]int64, 0)
	for _, lottery := range lotterys {
		lotteryIds = append(lotteryIds, lottery.Id)
	}
	if len(lotteryIds) == 0 {
		return &pb.GetLotteryPrizesListByUserIdResp{}, nil
	}
	//fmt.Println("lotteryIds", lotteryIds)

	// 根据lotteryIds获取奖品列表
	prizes, err := l.svcCtx.PrizeModel.FindAllByLotteryIds(l.ctx, lotteryIds)
	if err != nil {
		return nil, err
	}
	//for _, prize := range prizes {
	//	fmt.Println("prize", prize.Id)
	//}

	// 将奖品列表按照lotteryId进行分组
	prizesMap := make(map[int64][]*pb.Prize)
	for _, prize := range prizes {
		pbPrize := &pb.Prize{
			Id:         prize.Id,
			LotteryId:  prize.LotteryId,
			Name:       prize.Name,
			Level:      prize.Level,
			Thumb:      prize.Thumb,
			Count:      prize.Count,
			GrantType:  prize.GrantType,
			CreateTime: prize.CreateTime.Unix(),
			UpdateTime: prize.UpdateTime.Unix(),
		}
		prizesMap[prize.LotteryId] = append(prizesMap[prize.LotteryId], pbPrize)
	}
	//for k, v := range prizesMap {
	//	fmt.Println("prizesMap", k, v)
	//}
	//
	//// 组装数据
	resp := make([]*pb.LotteryPrizes, 0)
	for _, lottery := range lotterys {
		//fmt.Println("lottery", lottery.lotteryId)
		//fmt.Println("prizesMap", prizesMap[lottery.lotteryId])
		resp = append(resp, &pb.LotteryPrizes{
			LotteryId:       lottery.Id,
			Prizes:          prizesMap[lottery.Id],
			ParticipationId: lottery.ParticipationId,
			Time:            lottery.Time.Unix(),
		})
	}
	//for _, v := range resp {
	//	fmt.Println("resp", v)
	//}
	return &pb.GetLotteryPrizesListByUserIdResp{
		LotteryPrizes: resp,
	}, nil
}
