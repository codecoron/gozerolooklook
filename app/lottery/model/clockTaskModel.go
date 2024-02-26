package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClockTaskModel = (*customClockTaskModel)(nil)

type (
	// ClockTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClockTaskModel.
	ClockTaskModel interface {
		clockTaskModel
	}

	customClockTaskModel struct {
		*defaultClockTaskModel
	}
)

// NewClockTaskModel returns a model for the database table.
func NewClockTaskModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ClockTaskModel {
	return &customClockTaskModel{
		defaultClockTaskModel: newClockTaskModel(conn, c, opts...),
	}
}
