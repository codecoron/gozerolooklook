package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"
	"looklook/app/notice/cmd/rpc/notice"
	"looklook/common/constants"
	"looklook/common/xerr"
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

// PeopleLotteryStrategy 实现基于人数的抽奖策略
type PeopleLotteryStrategy struct {
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
func (l *AnnounceLotteryLogic) AnnounceLottery(in *pb.AnnounceLotteryReq) (*pb.AnnounceLotteryResp, error) {
	// 创建相应的抽奖策略
	var strategy LotteryStrategy
	switch in.AnnounceType {
	case constants.AnnounceTypeTimeLottery:
		// 开奖时间类型
		strategy = &TimeLotteryStrategy{
			AnnounceLotteryLogic: l,
			CurrentTime:          time.Now(),
		}
	case constants.AnnounceTypePeopleLottery:
		// 开奖时间类型
		strategy = &PeopleLotteryStrategy{
			AnnounceLotteryLogic: l,
			CurrentTime:          time.Now(),
		}
	}
	err := strategy.Run()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "AnnounceStrategy run error: %v", err)
	}
	//fmt.Println("AnnounceFinish") // t
	return &pb.AnnounceLotteryResp{}, nil
}

func (l *AnnounceLotteryLogic) DrawLottery(ctx context.Context, lotteryId int64, prizes []*model.Prize, participantor []int64) ([]Winner, error) {
	// test1： 用户有10个，奖品总数为5个，预计获奖人数945.即有某一时刻奖品数量为0。报错slice bounds out of range [7:2]    [已解决]
	// 随机选择中奖者
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// 获取奖品总数 = 中奖人数
	var WinnersNum int64
	for _, p := range prizes {
		WinnersNum += p.Count
	}

	winners := make([]Winner, 0)

	// 假定的每个用户的中奖倍率
	//testRatios := make([]int64, len(participantor))
	//for i := range testRatios {
	//	testRatios[i] = rand.Int63n(10) + 1 // Ensure a non-zero ratio, random value between 1 and 10
	//}

	records, err := l.svcCtx.ClockTaskRecordModel.GetClockTaskRecordByLotteryIdAndUserIds(lotteryId, participantor)
	if err != nil {
		return nil, err
	}

	//fmt.Println("records:", records)

	// 查出来可能有多条记录 每条记录就是完成的一次任务 increase_multiple就是那一次任务所增加的概率,一个用户可能有多条记录，我这边在业务里面再进行统计一次
	// 所以用一个map来存储每个用户的中奖倍率
	RationsMap := make(map[int64]int64)
	for _, participant := range participantor {
		RationsMap[participant] = 1
	}

	for _, record := range records {
		RationsMap[record.UserId] += record.IncreaseMultiple
	}

	Ratios := make([]int64, len(participantor))

	for i, participant := range participantor {
		Ratios[i] = RationsMap[participant]
	}

	//Ratios = testRatios

	//fmt.Println("Ratios:", Ratios)
	// 计算总的中奖概率
	totalRatio := int64(0)
	for _, ratio := range Ratios {
		totalRatio += ratio
	}
	// 计算每个用户的最终中奖概率
	FinalRatios := make([]float64, len(participantor))
	for idx := range Ratios {
		FinalRatios[idx] = float64(Ratios[idx]) / float64(totalRatio)
	}
	//fmt.Println("FinalRatios:", FinalRatios)

	// 根据中奖总数量进行开奖
	for i := 0; i < int(WinnersNum); i++ { // 中奖人数
		//fmt.Println("WinnersNum", i)
		var randomWinnerIndex int
		var winnerUserId int64

		//如果参与者少于预计中奖人数，就结束开奖。(参与人数 < 中奖人数)
		if len(participantor) == 0 {
			break
		}
		//else {
		//	// 随机选择一个参与者,得到中奖者的uid
		//	randomWinnerIndex = rand.Intn(len(participantor))
		//	winnerUserId = participantor[randomWinnerIndex]
		//}

		//生成一个0到1之间的随机数
		randomProbability := rand.Float64()

		// 根据随机数确定中奖用户
		probabilitySum := 0.0
		for idx := range participantor {
			// 逐个累加中奖概率，直到大于随机数
			probabilitySum += FinalRatios[idx]
			// 如果随机数小于等于累加的概率，说明中奖
			if randomProbability <= probabilitySum {
				// 中奖者的uid
				winnerUserId = participantor[idx]
				// 中奖者的索引
				randomWinnerIndex = idx
				break
			}
		}
		//fmt.Println("winnerUserId:", winnerUserId)
		//如果没有中奖用户，则第一个参与者中奖
		if winnerUserId == 0 {
			winnerUserId = participantor[0]
			//fmt.Println("没有中奖用户,默认第一个参与者中奖", winnerUserId)
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
		prizeId := prizes[0].Id

		// 创建中奖者对象
		winner := Winner{
			LotteryId: lotteryId,
			UserId:    winnerUserId,
			PrizeId:   prizeId, // 使用选中的奖品名称
		}

		winners = append(winners, winner)

		// 从参与者列表中移除已中奖的用户以及对应的中奖概率
		participantor = append(participantor[:randomWinnerIndex], participantor[randomWinnerIndex+1:]...)
		FinalRatios = append(FinalRatios[:randomWinnerIndex], FinalRatios[randomWinnerIndex+1:]...)
	}

	return winners, nil
}

// NotifyParticipators 通知MQ队列
func (l *AnnounceLotteryLogic) NotifyParticipators(participators []int64, lotteryId int64) error {
	fmt.Println("NotifyParticipators", participators, lotteryId)
	_, err := l.svcCtx.NoticeRpc.NoticeLotteryDraw(l.ctx, &notice.NoticeLotteryDrawReq{
		LotteryId: lotteryId,
		UserIds:   participators,
	})
	if err != nil {
		return err
	}
	return nil
}

// WriteWinnersToLotteryParticipation 更新参与抽奖表
func (l *AnnounceLotteryLogic) WriteWinnersToLotteryParticipation(winners []Winner) error {
	for _, w := range winners {
		err := l.svcCtx.LotteryParticipationModel.UpdateWinners(l.ctx, w.LotteryId, w.UserId, w.PrizeId)
		if err != nil {
			return err
		}
	}
	return nil
}

// Run 按时间开奖业务逻辑
func (s *TimeLotteryStrategy) Run() error {
	// 查询满足条件的抽奖
	lotteries, err := s.svcCtx.LotteryModel.GetLotterysByLessThanCurrentTime(s.ctx, s.CurrentTime, constants.AnnounceTypeTimeLottery)
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

			//fmt.Println("开始开奖的lottery:", lotteryId)

			// 获取该lotteryId对应的所有参与者
			participators, err = s.svcCtx.LotteryParticipationModel.GetParticipationUserIdsByLotteryId(s.ctx, lotteryId)
			if err != nil {
				return err
			}

			//testParticipators := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			//participators = testParticipators
			//fmt.Println("participators:", participators)

			winners, err := s.DrawLottery(s.ctx, lotteryId, prizes, participators)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "DrawLottery,lotteryId:%v,prizes:%v,participators:%v error: %v", lotteryId, prizes, participators, err)
			}

			// 测试查看所有winners
			//for _, w := range winners {
			//	fmt.Printf("testwinners:%+v\n", w)
			//}

			//更新抽奖状态为"已开奖"
			err = s.svcCtx.LotteryModel.UpdateLotteryStatus(s.ctx, lotteryId)
			if err != nil {
				return err
			}

			// 将得到的中奖信息，写入数据库participants
			err = s.WriteWinnersToLotteryParticipation(winners)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "AnnounceLotteryTrans error: %v", err)
		}

		// 执行开奖结果通知任务
		err := s.NotifyParticipators(participators, lotteryId)
		if err != nil {
			return err
		}
	}
	return err
}

// Run 按人数开奖策略
func (s *PeopleLotteryStrategy) Run() error {

	// 查询开奖类型为2并且没有开奖的所有抽奖
	lotteries, err := s.svcCtx.LotteryModel.GetTypeIs2AndIsNotAnnounceLotterys(s.ctx, constants.AnnounceTypePeopleLottery)
	if err != nil {
		return err
	}

	CheckedLottery, err := s.CheckLottery(lotteries)
	if err != nil {
		return err
	}
	// 遍历每一个抽奖
	for _, lottery := range CheckedLottery {
		var participators []int64
		// 事务处理
		err = s.svcCtx.LotteryModel.Trans(s.ctx, func(context context.Context, session sqlx.Session) error {
			//根据抽奖id得到对应的所有奖品
			prizes, err := s.svcCtx.PrizeModel.FindByLotteryId(s.ctx, lottery.Id)
			if err != nil {
				return err
			}

			//fmt.Println("开始开奖的lottery:", lottery.Id)

			// 获取该lotteryId对应的所有参与者
			participators, err = s.svcCtx.LotteryParticipationModel.GetParticipationUserIdsByLotteryId(s.ctx, lottery.Id)
			if err != nil {
				return err
			}

			//testParticipators := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

			//participators = testParticipators
			//fmt.Println("participators:", participators)

			winners, err := s.DrawLottery(s.ctx, lottery.Id, prizes, participators)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "DrawLottery,lotteryId:%v,prizes:%v,participators:%v, error: %v", lottery.Id, prizes, participators, err)
			}

			//测试查看所有winners
			//for _, w := range winners {
			//	fmt.Printf("testwinners:%+v\n", w)
			//}

			//更新抽奖状态为"已开奖"
			err = s.svcCtx.LotteryModel.UpdateLotteryStatus(s.ctx, lottery.Id)
			if err != nil {
				return err
			}

			// 将得到的中奖信息，写入数据库participants
			err = s.WriteWinnersToLotteryParticipation(winners)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.AnnounceLottery_ERROR), "AnnounceLotteryTrans error: %v", err)
		}

		// 执行开奖结果通知任务
		err := s.NotifyParticipators(participators, lottery.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *PeopleLotteryStrategy) CheckLottery(lotteries []*model.Lottery) (CheckedLotterys []*model.Lottery, err error) {
	// 筛选满足条件的抽奖
	// 1. 超过当前时间的，即使没有满足人数条件也需要进行开奖
	// 2. 当参与人数 >= 开奖人数，进行开奖

	for _, l := range lotteries {
		// l.AnnounceTime 小于等于 s.CurrentTime,即超过当前时间，需要开奖
		if l.AnnounceTime.Before(s.CurrentTime) || l.AnnounceTime.Equal(s.CurrentTime) {
			CheckedLotterys = append(CheckedLotterys, l)
		} else {
			//fmt.Println("lotteryId:", l.Id)
			ParticipatorCount, err := s.svcCtx.LotteryParticipationModel.GetParticipatorsCountByLotteryId(s.ctx, l.Id)
			//fmt.Println("ParticipatorCount:", ParticipatorCount)
			if err != nil {
				return nil, err
			}
			// 检查参与人数是否达到指定人数
			if ParticipatorCount >= l.JoinNumber {
				CheckedLotterys = append(CheckedLotterys, l)
			}
		}
	}
	return
}

// 在开奖模块中，采用定时任务的方式，每隔一段时间执行一次开奖任务，这样可以避免在高并发情况下，大量的开奖任务同时执行，导致数据库压力过大。
// 为什么不一直监听MQ消息，而是采用定时任务的方式？
// 1. MQ消息监听方式，需要一直监听MQ消息，这样会导致大量的MQ消息监听，对MQ服务器压力过大。

// 延迟队列会被问到什么面试题？
// 1. 延迟队列是什么？
// 2. 延迟队列的实现原理是什么？
// 3. 延迟队列的应用场景有哪些？
// 4. 延迟队列的实现方式有哪些？
// 5. 延迟队列的使用有哪些注意事项？
// 请你给出上面问题的答案。
// 1. 延迟队列是一种特殊的消息队列，它的消息不会立即被消费，而是在一定时间后才会被消费。
// 2. 延迟队列的实现原理是通过消息的TTL（Time To Live）和死信队列来实现的。消息的TTL是指消息的存活时间，当消息的TTL到期后，消息会被发送到死信队列中，然后再由消费者来消费。
// 3. 延迟队列的应用场景有很多，比如订单超时未支付，可以将订单消息发送到延迟队列中，然后在一定时间后再进行处理；还有比如秒杀活动，可以将秒杀消息发送到延迟队列中，然后在活动开始后再进行处理。
// 4. 延迟队列的实现方式有很多，比如可以通过消息队列的TTL和死信队列来实现，也可以通过定时任务来实现，还可以通过定时轮询来实现。
// asynq 是一个简单、可靠、高效的分布式任务队列，它支持延迟任务、重试任务、定时任务、并发任务等功能，适用于各种异步任务处理场景。
// asynq的实现原理是通过消息队列来实现的，它使用了Redis作为消息队列，通过Redis的list数据结构来存储消息，通过Redis的pub/sub功能来实现消息的发布和订阅，通过Redis的zset数据结构来实现消息的延迟和定时。
// asynq的应用场景有很多，比如可以用来处理异步任务、定时任务、延迟任务、重试任务等，适用于各种异步任务处理场景。
// asynq的实现方式是通过Redis的list数据结构来存储消息，通过Redis的pub/sub功能来实现消息的发布和订阅，通过Redis的zset数据结构来实现消息的延迟和定时。
// asynq的使用有一些注意事项，比如需要保证消息的幂等性，需要保证消息的可靠性，需要保证消息的顺序性，需要保证消息的一致性等。
