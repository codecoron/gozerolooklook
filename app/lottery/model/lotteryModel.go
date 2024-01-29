package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"math/rand"
	"time"
)

var _ LotteryModel = (*customLotteryModel)(nil)

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		// 自定义方法
		UpdatePublishTime(ctx context.Context, data *Lottery) error
		LotteryList(ctx context.Context, page, limit, selected, lastId int64) ([]*Lottery, error)
		FindUserIdByLotteryId(ctx context.Context, lotteryId int64) (*int64, error)
		QueryLotteries(ctx context.Context, currentTime time.Time) ([]int64, error)
		DrawLottery(ctx context.Context, lottery *Lottery, prizes, zeroPrizes []*Prize) ([]Winner, error)
		WriteResultToDB(ctx context.Context, winners []*Winner) error
		UpdateLotteryStatus(ctx context.Context, lotteryID int64) error
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}
)

func (m *defaultLotteryModel) UpdatePublishTime(ctx context.Context, data *Lottery) error {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set publish_time = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, data.PublishTime, data.Id)
	}, lotteryLotteryIdKey)
	return err
}

func (c *customLotteryModel) LotteryList(ctx context.Context, page, limit, selected, lastId int64) ([]*Lottery, error) {
	var query string
	if selected != 0 {
		query = fmt.Sprintf("select %s from %s where id > ? and is_selected = 1 limit ?,?", lotteryRows, c.table)
	} else {
		query = fmt.Sprintf("select %s from %s where id > ? limit ?,?", lotteryRows, c.table)
	}
	var resp []*Lottery
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lastId, (page-1)*limit, limit)
	return resp, err
}

func (c *customLotteryModel) FindUserIdByLotteryId(ctx context.Context, lotteryId int64) (*int64, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, lotteryId)
	var resp int64
	err := c.QueryRowCtx(ctx, &resp, lotteryLotteryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select user_id from %s where id = ?", c.table)
		return conn.QueryRowCtx(ctx, v, query, lotteryId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c, opts...),
	}
}

func (c *customLotteryModel) QueryLotteries(ctx context.Context, currentTime time.Time) ([]int64, error) {
	var resp []int64
	query := fmt.Sprintf("SELECT id FROM %s WHERE announce_type = 1 AND is_announced = 0 AND announce_time <= ?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, currentTime)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type Winner struct {
	LotteryId int64
	UserId    int64
	PrizeId   int64
}

func (c *customLotteryModel) DrawLottery(ctx context.Context, lottery *Lottery, prizes, zeroPrizes []*Prize) ([]Winner, error) {
	//var participators []int64
	//query := fmt.Sprintf("SELECT user_id FROM lottery_participants WHERE lottery_id = ?")
	//err := c.QueryRowsNoCacheCtx(ctx, &participants, query, lottery.Id)
	//if err != nil {
	//	return nil, err
	//}

	testParticipators := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

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
		if len(testParticipators) == 0 {
			//fmt.Println("如果参与者少于预计中奖人数，就结束开奖。")
			break
		} else {
			// 随机选择一个参与者,得到中奖者的uid
			randomWinnerIndex = rand.Intn(len(testParticipators))
			winnerUserId = testParticipators[randomWinnerIndex]
		}

		// 随机选择一个奖品，确保奖品的剩余数量大于0
		prize := new(Prize)
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
				prizes = append(prizes[:randomPrizeIndex], prizes[randomPrizeIndex+1:]...)
			}
		}
		// 正常情况是参与人数 <>= 预计参与人数; 预计参与人数 <= 奖品总数量

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
		testParticipators = append(testParticipators[:randomWinnerIndex], testParticipators[randomWinnerIndex+1:]...)

		// 减少选中奖品的剩余数量；如果当前prize的数量为0，则去掉这个奖品
		prizes[randomPrizeIndex].Count--
		if prizes[randomPrizeIndex].Count == 0 {
			zeroPrizes = append(zeroPrizes, prizes[randomPrizeIndex])
			prizes = append(prizes[:randomPrizeIndex], prizes[randomPrizeIndex+1:]...)
		}
		for _, p := range prizes {
			fmt.Printf("%+v\n", p)
		}
	}

	return winners, nil
}

// WriteResultToDB todo 将resultList换成表对应的数据结构
func (c *customLotteryModel) WriteResultToDB(ctx context.Context, winners []*Winner) error {
	// 准备插入数据的SQL语句
	query := "INSERT INTO lottery_result (result) VALUES (?)"

	err := c.Trans(ctx, func(ctx context.Context, session sqlx.Session) error {
		// 执行插入操作
		for _, result := range winners {
			_, err := session.Exec(query, result)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateLotteryStatus 根据lotteryId更新lottery状态为已开奖
func (c *customLotteryModel) UpdateLotteryStatus(ctx context.Context, lotteryID int64) error {
	// 准备更新数据的SQL语句
	query := fmt.Sprintf("UPDATE %s SET is_announced = 1 WHERE id = ?", c.table)

	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.Exec(query, lotteryID)
	})
	if err != nil {
		return err
	}
	return nil
}
