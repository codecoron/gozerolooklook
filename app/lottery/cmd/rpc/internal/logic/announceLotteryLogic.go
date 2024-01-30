package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"
	"looklook/app/notice/cmd/rpc/notice"
	"math/rand"
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
		//fmt.Println("testttasdf", err) // f
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

			//todo 优化代码
			// 根据lotteryId获取一个lottery详情，需要joinNumber字段 todo 优化名字 按人数开奖的指定人数
			lottery, err := s.svcCtx.LotteryModel.FindOne(s.ctx, lotteryId)
			if err != nil {
				return err
			}
			fmt.Println("开始开奖的lottery:", lottery.Id)

			// 开奖，根据抽奖规则以及参与抽奖用户表，得到最终的获奖名单winners
			// 利用go切片的特性，传入cpPrizes，他们都指向同一片地址，在DrawLottery内部对cpPrizes的修改也是对Prizes的修改，但是cpPrizes不论是删除元素还是扩容都不影响Prizes底层的数组，符合业务逻辑
			zeroPrizes := make([]*model.Prize, 0)

			// todo 获取该lotteryId对应的所有参与者
			//var participators []int64
			//query := fmt.Sprintf("SELECT user_id FROM lottery_participants WHERE lottery_id = ?")
			//err := c.QueryRowsNoCacheCtx(ctx, &participants, query, lottery.Id)
			//if err != nil {
			//	return nil, err
			//}

			testParticipators := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			participators = testParticipators

			winners, err := s.DrawLottery(s.ctx, lottery, prizes, zeroPrizes, participators)
			if err != nil {
				return err
			}

			// 测试查看所有winners
			for _, w := range winners {
				fmt.Printf("testwinners:%+v\n", w)
			}

			//更新抽奖状态为"已开奖" t
			err = s.svcCtx.LotteryModel.UpdateLotteryStatus(s.ctx, lottery.Id)
			if err != nil {
				return err
			}

			// 更新数据库中Prize表该奖品的数量
			for _, p := range prizes {
				fmt.Println("prizeId:", p.Id, "prizeCount:", p.Count)
				err = s.svcCtx.PrizeModel.Update(s.ctx, p)
				if err != nil {
					return err
				}
			}

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

		// TODO 执行开奖结果通知任务
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
func (s *TimeLotteryStrategy) DrawLottery(ctx context.Context, lottery *model.Lottery, prizes, zeroPrizes []*model.Prize, participantor []int64) ([]Winner, error) {
	// test1： 用户有10个，奖品总数为5个，预计获奖人数945.即有某一时刻奖品数量为0。报错slice bounds out of range [7:2]    [已解决]
	// 随机选择中奖者
	rand.New(rand.NewSource(time.Now().UnixNano()))
	WinnersNum := lottery.JoinNumber // 中奖者个数,joinNumber在按时开奖的时候表示中奖者个数
	winners := make([]Winner, 0)

	//fmt.Println("alltestNumber", testParticipators)
	for i := 0; i < int(WinnersNum); i++ {
		//fmt.Println("WinnersNum", i)
		var randomWinnerIndex int
		var winnerUserId int64

		// 如果参与者少于预计中奖人数，就结束开奖。(参与人数 < 预计中奖人数 && 奖品数量 >= 参与人数)
		if len(participantor) == 0 {
			//fmt.Println("如果参与者少于预计中奖人数，就结束开奖。")
			break
		} else {
			// 随机选择一个参与者,得到中奖者的uid
			randomWinnerIndex = rand.Intn(len(participantor))
			winnerUserId = participantor[randomWinnerIndex]
		}

		// 随机选择一个奖品，确保奖品的剩余数量大于0
		prize := new(model.Prize)
		var randomPrizeIndex int
		for {
			// 如果所有奖品的剩余数量都为0，直接返回即可。(该情况可能发生的情形为预计中奖人数 >= 参与人数 > 奖品数量)
			if len(prizes) == 0 {
				fmt.Println("no prizes left")
				return winners, nil
			}
			randomPrizeIndex = rand.Intn(len(prizes))
			prize = prizes[randomPrizeIndex]
			if prize.Count > 0 {
				break
			} else {
				//
				prizes = append(prizes[:randomPrizeIndex], prizes[randomPrizeIndex+1:]...)
			}
		}
		// 正常情况是参与人数 <>= 奖品数量; 预计中奖人数 == 奖品总数量

		// 创建中奖者对象
		winner := Winner{
			LotteryId: lottery.Id,
			UserId:    winnerUserId,
			PrizeId:   prize.Id, // 使用选中的奖品名称
		}

		//fmt.Printf("%+v\n", winner)
		//fmt.Println("lenOfPrizes:", len(prizes), "lenOfTestNumber:", len(testParticipators))

		winners = append(winners, winner)

		// 从参与者列表中移除已中奖的用户
		participantor = append(participantor[:randomWinnerIndex], participantor[randomWinnerIndex+1:]...)

		// 减少选中奖品的剩余数量；如果当前prize的数量为0，则去掉这个奖品
		prizes[randomPrizeIndex].Count--
		if prizes[randomPrizeIndex].Count == 0 {
			zeroPrizes = append(zeroPrizes, prizes[randomPrizeIndex])
			prizes = append(prizes[:randomPrizeIndex], prizes[randomPrizeIndex+1:]...)
		}
		// 测试
		for _, p := range prizes {
			fmt.Printf("%+v\n", p)
		}
	}

	return winners, nil
}
