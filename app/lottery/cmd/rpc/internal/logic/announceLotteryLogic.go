package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"
	"looklook/app/notice/cmd/rpc/notice"
	"math/rand"
	"sort"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/lottery/cmd/rpc/internal/svc"
)

type AnnounceLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// LotteryStrategy 定义抽奖策略接口
type LotteryStrategy interface {
	Run() error
}

// TimeLotteryStrategy 实现基于时间的抽奖策略
type TimeLotteryStrategy struct {
	*AnnounceLotteryLogic
	CurrentTime time.Time
}

type Winner struct {
	LotteryId int64
	UserId    int64
	PrizeId   int64
}

func NewAnnounceLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnnounceLotteryLogic {
	return &AnnounceLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AnnounceLotteryLogic) AnnounceLottery(in *pb.AnnounceLotteryReq) (*pb.AnnounceLotteryResp, error) {
	// 创建相应的抽奖策略
	var strategy LotteryStrategy
	switch in.AnnounceType {
	case 1:
		// 开奖时间类型
		strategy = &TimeLotteryStrategy{
			AnnounceLotteryLogic: l,
			CurrentTime:          time.Now(),
		}
	}
	err := strategy.Run()
	if err != nil {
		return nil, err
	}
	fmt.Println("AnnounceFinish") // t
	return &pb.AnnounceLotteryResp{}, nil
}

// NotifyParticipators 通知MQ队列
func (l *AnnounceLotteryLogic) NotifyParticipators(participators []int64, lotteryId int64) error {
	_, err := l.svcCtx.NoticeRpc.NoticeLotteryDraw(l.ctx, &notice.NoticeLotteryDrawReq{
		LotteryId: lotteryId,
		UserIds:   participators,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *TimeLotteryStrategy) Run() error {
	// 查询满足条件的抽奖
	lotteries, err := s.svcCtx.LotteryModel.QueryLotteries(s.ctx, s.CurrentTime)
	if err != nil {
		return err
	}

	// 遍历每一个抽奖
	for _, lotteryId := range lotteries {
		var participators []int64
		// 事务处理
		err = s.svcCtx.LotteryModel.Trans(s.ctx, func(context context.Context, session sqlx.Session) error {
			//根据抽奖id得到对应的所有奖品
			prizes, err := s.svcCtx.PrizeModel.FindByLotteryId(s.ctx, lotteryId)
			if err != nil {
				return err
			}

			fmt.Println("开始开奖的lottery:", lotteryId)

			// todo 获取该lotteryId对应的所有参与者
			//var participators []int64
			//query := fmt.Sprintf("SELECT user_id FROM lottery_participants WHERE lottery_id = ?")
			//err := c.QueryRowsNoCacheCtx(ctx, &participants, query, lottery.Id)
			//if err != nil {
			//	return nil, err
			//}

			testParticipators := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			participators = testParticipators

			winners, err := s.DrawLottery(s.ctx, lotteryId, prizes, participators)
			if err != nil {
				return err
			}

			// 测试查看所有winners
			for _, w := range winners {
				fmt.Printf("testwinners:%+v\n", w)
			}

			//更新抽奖状态为"已开奖" t
			//err = s.svcCtx.LotteryModel.UpdateLotteryStatus(s.ctx, lotteryId)
			//if err != nil {
			//	return err
			//}

			// 更新数据库中Prize表该奖品的数量
			//for _, p := range prizes {
			//	fmt.Println("prizeId:", p.Id, "prizeCount:", p.Count)
			//	err = s.svcCtx.PrizeModel.Update(s.ctx, p)
			//	if err != nil {
			//		return err
			//	}
			//}

			// todo 将得到的中奖信息，写入数据库participants
			//err = s.svcCtx.LotteryModel.WriteResultToDB(s.ctx, winners)
			//if err != nil {
			//	return err
			//}
			return nil
		})
		if err != nil {
			return err
		}

		// 执行开奖结果通知任务
		err := s.NotifyParticipators(participators, lotteryId)
		if err != nil {
			return err
		}
	}
	return err
}

/*
*
开奖：
1 按顺序开奖 一等奖1个 二等经2个  2个人参与
2 伪代码：
获得中奖人数 m
获得奖品列表 奖品level升序 总数n
m<=n
中奖人数n和奖品m做匹配，n1 = m1
map[userid]{prize}

	for m {
		for n{
			m1 = n1
		}
	}
*/

func (s *TimeLotteryStrategy) DrawLottery(ctx context.Context, lotteryId int64, prizes []*model.Prize, participantor []int64) ([]Winner, error) {
	// test1： 用户有10个，奖品总数为5个，预计获奖人数945.即有某一时刻奖品数量为0。报错slice bounds out of range [7:2]    [已解决]
	// 随机选择中奖者
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 获取奖品总数 = 中奖人数
	var WinnersNum int64
	for _, p := range prizes {
		WinnersNum += p.Count
	}

	winners := make([]Winner, 0)

	for i := 0; i < int(WinnersNum); i++ { // 中奖人数
		//fmt.Println("WinnersNum", i)
		var randomWinnerIndex int
		var winnerUserId int64

		// 如果参与者少于预计中奖人数，就结束开奖。(参与人数 < 中奖人数)
		if len(participantor) == 0 {
			break
		} else {
			// 随机选择一个参与者,得到中奖者的uid
			randomWinnerIndex = rand.Intn(len(participantor))
			winnerUserId = participantor[randomWinnerIndex]
		}

		// 对所有prizes按照type排序 // todo 获取的时候能保证type有序吗？有序则可以不用排序了
		sort.Slice(prizes, func(i, j int) bool {
			return prizes[i].Type < prizes[j].Type
		})

		// 如果当前等级的奖品的剩余数量都为0，去掉，获取下一等级的奖品。
		if prizes[0].Count == 0 {
			prizes = prizes[1:]
		}
		prizes[0].Count--
		//fmt.Printf("testPrize:%+v\n", prizes[0])
		prizeId := prizes[0].Id

		// 创建中奖者对象
		winner := Winner{
			LotteryId: lotteryId,
			UserId:    winnerUserId,
			PrizeId:   prizeId, // 使用选中的奖品名称
		}

		winners = append(winners, winner)

		// 从参与者列表中移除已中奖的用户
		participantor = append(participantor[:randomWinnerIndex], participantor[randomWinnerIndex+1:]...)
	}

	return winners, nil
}
