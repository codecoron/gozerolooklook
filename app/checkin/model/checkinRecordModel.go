package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CheckinRecordModel = (*customCheckinRecordModel)(nil)

type (
	// CheckinRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCheckinRecordModel.
	CheckinRecordModel interface {
		checkinRecordModel
		// 自定义方法
		FindOneByUserId(ctx context.Context, userId int64) (*CheckinRecord, error)
		UpdateByUserId(ctx context.Context, data *CheckinRecord) error
	}

	customCheckinRecordModel struct {
		*defaultCheckinRecordModel
	}
)

func (m *defaultCheckinRecordModel) UpdateByUserId(ctx context.Context, data *CheckinRecord) error {
	checkinCheckinRecordUserIdKey := fmt.Sprintf("%s%v", "cache:checkin:checkinRecord:userId:", data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `user_id` = ?", m.table, checkinRecordRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ContinuousCheckinDays, data.State, data.LastCheckinDate, data.UserId)
	}, checkinCheckinRecordUserIdKey)
	return err
}

func (m *defaultCheckinRecordModel) FindOneByUserId(ctx context.Context, userId int64) (*CheckinRecord, error) {
	checkinCheckinRecordUserIdKey := fmt.Sprintf("%s%v", "cache:checkin:checkinRecord:userId:", userId)
	var resp CheckinRecord
	err := m.QueryRowCtx(ctx, &resp, checkinCheckinRecordUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", checkinRecordRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userId)
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

// NewCheckinRecordModel returns a model for the database table.
func NewCheckinRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CheckinRecordModel {
	return &customCheckinRecordModel{
		defaultCheckinRecordModel: newCheckinRecordModel(conn, c, opts...),
	}
}
