package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClockTaskRecordModel = (*customClockTaskRecordModel)(nil)

type (
	// ClockTaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClockTaskRecordModel.
	ClockTaskRecordModel interface {
		clockTaskRecordModel
	}

	customClockTaskRecordModel struct {
		*defaultClockTaskRecordModel
	}
)

// NewClockTaskRecordModel returns a model for the database table.
func NewClockTaskRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ClockTaskRecordModel {
	return &customClockTaskRecordModel{
		defaultClockTaskRecordModel: newClockTaskRecordModel(conn, c, opts...),
	}
}
