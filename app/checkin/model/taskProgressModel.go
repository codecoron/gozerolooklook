package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskProgressModel = (*customTaskProgressModel)(nil)

type (
	// TaskProgressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskProgressModel.
	TaskProgressModel interface {
		taskProgressModel
	}

	customTaskProgressModel struct {
		*defaultTaskProgressModel
	}
)

// NewTaskProgressModel returns a model for the database table.
func NewTaskProgressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskProgressModel {
	return &customTaskProgressModel{
		defaultTaskProgressModel: newTaskProgressModel(conn, c, opts...),
	}
}
