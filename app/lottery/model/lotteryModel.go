package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
