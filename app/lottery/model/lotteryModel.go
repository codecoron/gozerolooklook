package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LotteryModel = (*customLotteryModel)(nil)

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		LotteryList(ctx context.Context, page, limit, selected, lastId int64) ([]*Lottery, error)
		FindUserIdByLotteryId(ctx context.Context, lotteryId int64) (*int64, error)
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
