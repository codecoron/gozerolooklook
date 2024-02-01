package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TaskRecordModel = (*customTaskRecordModel)(nil)

type (
	// TaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskRecordModel.
	TaskRecordModel interface {
		taskRecordModel
		FindByUserId(ctx context.Context, userId int64, builder squirrel.SelectBuilder, orderBy string) ([]*TaskRecord, error)
	}

	customTaskRecordModel struct {
		*defaultTaskRecordModel
	}
)

func (m *defaultTaskRecordModel) FindByUserId(ctx context.Context, userId int64, builder squirrel.SelectBuilder, orderBy string) ([]*TaskRecord, error) {

	builder = builder.Columns(taskRecordRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("user_id = ?", userId).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TaskRecord
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// NewTaskRecordModel returns a model for the database table.
func NewTaskRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskRecordModel {
	return &customTaskRecordModel{
		defaultTaskRecordModel: newTaskRecordModel(conn, c, opts...),
	}
}
