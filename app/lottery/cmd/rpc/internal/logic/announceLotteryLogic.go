package logic

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/lottery/cmd/rpc/internal/svc"
)

type AnnounceLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnnounceLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnnounceLotteryLogic {
	return &AnnounceLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
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

func (s *TimeLotteryStrategy) Run() error {
	// 查询满足条件的抽奖
	lotteries, err := s.svcCtx.LotteryModel.QueryLotteries(s.ctx, s.CurrentTime)
	if err != nil {
		//fmt.Println("testttasdf", err) // f
		return err
	}

	// 遍历每一个抽奖
	for _, lotteryId := range lotteries {
		// 事务处理
		err = s.svcCtx.LotteryModel.Trans(s.ctx, func(context context.Context, session sqlx.Session) error {
			//根据抽奖id得到对应的所有奖品
			prizes, err := s.svcCtx.PrizeModel.FindByLotteryId(s.ctx, lotteryId)
			if err != nil {
				return err
			}

			// 根据lotteryId获取一个lottery详情，需要joinNumber字段
			lottery, err := s.svcCtx.LotteryModel.FindOne(s.ctx, lotteryId)
			if err != nil {
				return err
			}
			fmt.Println("开始开奖的lottery:", lottery.Id)

			// 开奖，根据抽奖规则以及参与抽奖用户表，得到最终的获奖名单winners
			// 利用go切片的特性，传入cpPrizes，他们都指向同一片地址，在DrawLottery内部对cpPrizes的修改也是对Prizes的修改，但是cpPrizes不论是删除元素还是扩容都不影响Prizes底层的数组，符合业务逻辑
			zeroPrizes := make([]*model.Prize, 0)
			winners, err := s.svcCtx.LotteryModel.DrawLottery(s.ctx, lottery, prizes, zeroPrizes)
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

			// todo 将得到的中奖信息，通过WriteResultToDB函数写入数据库
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
		//notifyWinners(winners)
	}
	return err
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

func notifyWinners(winners []*model.Winner) {
	// 执行开奖结果通知任务，可以是发送通知消息或邮件等操作
	// ...
	// TODO 通知MQ队列
}
