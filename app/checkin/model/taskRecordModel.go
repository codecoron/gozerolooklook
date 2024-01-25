package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskRecordModel = (*customTaskRecordModel)(nil)

type (
	// TaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskRecordModel.
	TaskRecordModel interface {
		taskRecordModel
	}

	customTaskRecordModel struct {
		*defaultTaskRecordModel
	}
)

// NewTaskRecordModel returns a model for the database table.
func NewTaskRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskRecordModel {
	return &customTaskRecordModel{
		defaultTaskRecordModel: newTaskRecordModel(conn, c, opts...),
	}
}
