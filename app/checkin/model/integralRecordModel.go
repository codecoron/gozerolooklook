package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ IntegralRecordModel = (*customIntegralRecordModel)(nil)

type (
	// IntegralRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customIntegralRecordModel.
	IntegralRecordModel interface {
		integralRecordModel
	}

	customIntegralRecordModel struct {
		*defaultIntegralRecordModel
	}
)

// NewIntegralRecordModel returns a model for the database table.
func NewIntegralRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) IntegralRecordModel {
	return &customIntegralRecordModel{
		defaultIntegralRecordModel: newIntegralRecordModel(conn, c, opts...),
	}
}
