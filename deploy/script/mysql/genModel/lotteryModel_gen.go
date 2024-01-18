// Code generated by goctl. DO NOT EDIT.

package genModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	lotteryFieldNames          = builder.RawFieldNames(&Lottery{})
	lotteryRows                = strings.Join(lotteryFieldNames, ",")
	lotteryRowsExpectAutoSet   = strings.Join(stringx.Remove(lotteryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	lotteryRowsWithPlaceHolder = strings.Join(stringx.Remove(lotteryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLotteryLotteryIdPrefix = "cache:lottery:lottery:id:"
)

type (
	lotteryModel interface {
		Insert(ctx context.Context, data *Lottery) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Lottery, error)
		Update(ctx context.Context, data *Lottery) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLotteryModel struct {
		sqlc.CachedConn
		table string
	}

	Lottery struct {
		Id            int64        `db:"id"`
		UserId        int64        `db:"user_id"`        // 发起抽奖用户ID
		Name          string       `db:"name"`           // 默认取一等奖名称
		Thumb         string       `db:"thumb"`          // 默认取一等经配图
		PublishType   int64        `db:"publish_type"`   // 开奖设置：1按时间开奖 2按人数开奖 3即抽即中
		PublishTime   sql.NullTime `db:"publish_time"`   // 开奖时间
		JoinNumber    int64        `db:"join_number"`    // 自动开奖人数
		Introduce     string       `db:"introduce"`      // 抽奖说明
		AwardDeadline time.Time    `db:"award_deadline"` // 领奖截止时间
		CreateTime    time.Time    `db:"create_time"`
		UpdateTime    time.Time    `db:"update_time"`
		IsSelected    int64        `db:"is_selected"` // 是否精选 1是 0否
	}
)

func newLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultLotteryModel {
	return &defaultLotteryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`lottery`",
	}
}

func (m *defaultLotteryModel) Delete(ctx context.Context, id int64) error {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, lotteryLotteryIdKey)
	return err
}

func (m *defaultLotteryModel) FindOne(ctx context.Context, id int64) (*Lottery, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, id)
	var resp Lottery
	err := m.QueryRowCtx(ctx, &resp, lotteryLotteryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", lotteryRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultLotteryModel) Insert(ctx context.Context, data *Lottery) (sql.Result, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, lotteryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Name, data.Thumb, data.PublishType, data.PublishTime, data.JoinNumber, data.Introduce, data.AwardDeadline, data.IsSelected)
	}, lotteryLotteryIdKey)
	return ret, err
}

func (m *defaultLotteryModel) Update(ctx context.Context, data *Lottery) error {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, lotteryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Name, data.Thumb, data.PublishType, data.PublishTime, data.JoinNumber, data.Introduce, data.AwardDeadline, data.IsSelected, data.Id)
	}, lotteryLotteryIdKey)
	return err
}

func (m *defaultLotteryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, primary)
}

func (m *defaultLotteryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", lotteryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLotteryModel) tableName() string {
	return m.table
}
