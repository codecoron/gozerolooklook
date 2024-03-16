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
	// 获取当前用户发起的所有抽奖
	lotterys, err := l.svcCtx.LotteryModel.FindAllByUserId(in.UserId)
	if err != nil {
		return nil, err
	}
	//for _, lottery := range lotterys {
	//	fmt.Println("获取当前用户发起的所有抽奖", lottery.CreateTime)
	//}
	// 获取当前用户所有参与的抽奖信息
	participates, err := l.svcCtx.LotteryParticipationModel.FindAllByUserId(in.UserId)
	if err != nil {
		return nil, err
	}
	//for _, participation := range participates {
	//	fmt.Println("获取当前用户所有参与的抽奖信息", participation.CreateTime)
	//	fmt.Println("获取当前用户所有参与的fff", participation.UpdateTime)
	//}
	LotteryList := make([]*LotteryPrizes, 0)
	if in.Type == 1 {
		// 获取当前用户所有参与+发起的抽奖信息
		alreadyExist := make(map[int64]struct{})
		for _, participation := range participates {
			LotteryList = append(LotteryList, &LotteryPrizes{
				lotteryId: participation.LotteryId,
			})
			alreadyExist[participation.LotteryId] = struct{}{}
		}
		// 去重

		for _, lottery := range lotterys {
			if _, ok := alreadyExist[lottery.Id]; !ok {
				LotteryList = append(LotteryList, &LotteryPrizes{
					lotteryId: lottery.Id,
				})
			}
		}
	} else if in.Type == 2 {
		// 获取当前用户发起的所有抽奖
		for _, participation := range participates {
			LotteryList = append(LotteryList, &LotteryPrizes{
				lotteryId:  participation.LotteryId,
				CreateTime: participation.CreateTime.Unix(),
			})
		}
	} else if in.Type == 3 {
		// 获取当前用户中奖的所有抽奖
		for _, participation := range participates {
			if participation.IsWon == 1 {
				LotteryList = append(LotteryList, &LotteryPrizes{
					lotteryId: participation.LotteryId,
					WonTime:   participation.UpdateTime.Unix(),
				})
			}
		}
	}

	if len(LotteryList) == 0 {
		return &pb.GetLotteryPrizesListByUserIdResp{}, nil
	}
	// 根据page和size获取当前用户所有的抽奖记录
	// 处理边界问题
	if int((in.Page-1)*in.Size) > len(LotteryList) {
		return &pb.GetLotteryPrizesListByUserIdResp{}, nil
	}

	if in.Page > 0 && in.Size > 0 && int(in.Page*in.Size) > len(LotteryList) {
		LotteryList = LotteryList[(in.Page-1)*in.Size:]
	} else if in.Page > 0 && in.Size > 0 {
		LotteryList = LotteryList[(in.Page-1)*in.Size : in.Page*in.Size]
	} else {
		return &pb.GetLotteryPrizesListByUserIdResp{}, nil
	}

	// 得到LottoryList后，根据lotteryId获取奖品列表
	// 先统计lotteryIds
	lotteryIds := make([]int64, 0)
	for _, lottery := range LotteryList {
		lotteryIds = append(lotteryIds, lottery.lotteryId)
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

	// 组装数据
	resp := make([]*pb.LotteryPrizes, 0)
	for _, lottery := range LotteryList {
		//fmt.Println("lottery", lottery.lotteryId)
		//fmt.Println("prizesMap", prizesMap[lottery.lotteryId])
		resp = append(resp, &pb.LotteryPrizes{
			LotteryId:       lottery.lotteryId,
			Prizes:          prizesMap[lottery.lotteryId],
			ParticipateTime: lottery.ParticipateTime,
			CreateTime:      lottery.CreateTime,
			WonTime:         lottery.WonTime,
		})
	}
	//for _, v := range resp {
	//	fmt.Println("resp", v)
	//}
	return &pb.GetLotteryPrizesListByUserIdResp{
		LotteryPrizes: resp,
	}, nil
}
