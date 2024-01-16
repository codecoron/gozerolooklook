package model

import (
	"context"
	"database/sql"
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
		//todo 自定义的方法写到这里，避免被覆盖掉 ; 重写模板 提高效率
		List(ctx context.Context, page, limit int64) ([]*Lottery, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Lottery) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}
)

func (c *customLotteryModel) List(ctx context.Context, page, limit int64) ([]*Lottery, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", lotteryRows, c.table)
	var resp []*Lottery
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *customLotteryModel) TransInsert(ctx context.Context, session sqlx.Session, data *Lottery) (sql.Result, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, lotteryRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.Name, data.Thumb, data.PublishType, data.PublishTime, data.JoinNumber, data.Introduce, data.AwardDeadline, data.IsSelected)
	}, lotteryLotteryIdKey)
	return ret, err
}

func (m *customLotteryModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c, opts...),
	}
}
