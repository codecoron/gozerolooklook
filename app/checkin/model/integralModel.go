package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IntegralModel = (*customIntegralModel)(nil)

type (
	// IntegralModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIntegralModel.
	IntegralModel interface {
		integralModel
		FindOneByUserId(ctx context.Context, userId int64) (*Integral, error)
		UpdateByUserId(ctx context.Context, data *Integral) error
	}

	customIntegralModel struct {
		*defaultIntegralModel
	}
)

func (m *defaultIntegralModel) UpdateByUserId(ctx context.Context, data *Integral) error {
	checkinIntegralUserIdKey := fmt.Sprintf("%s%v", "cache:checkin:integral:userId:", data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, integralRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Integral, data.UserId)
	}, checkinIntegralUserIdKey)
	return err
}

func (m *defaultIntegralModel) FindOneByUserId(ctx context.Context, userId int64) (*Integral, error) {
	checkinIntegralUserIdKey := fmt.Sprintf("%s%v", "cache:checkin:integral:userId:", userId)
	var resp Integral
	err := m.QueryRowCtx(ctx, &resp, checkinIntegralUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", integralRows, m.table)
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

// NewIntegralModel returns a model for the database table.
func NewIntegralModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IntegralModel {
	return &customIntegralModel{
		defaultIntegralModel: newIntegralModel(conn, c, opts...),
	}
}
