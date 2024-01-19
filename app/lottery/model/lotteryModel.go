package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LotteryModel = (*customLotteryModel)(nil)

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		LotteryList(ctx context.Context, page, limit, selected, lastId int64) ([]*Lottery, error)
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}
)

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

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c, opts...),
	}
}
