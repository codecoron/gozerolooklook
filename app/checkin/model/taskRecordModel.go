package model

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ TaskRecordModel = (*customTaskRecordModel)(nil)

type (
	// TaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskRecordModel.
	TaskRecordModel interface {
		taskRecordModel
		FindByUserId(ctx context.Context, userId int64, builder squirrel.SelectBuilder, orderBy string) ([]*TaskRecord, error)
		FindByUserIdAndTaskId(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error)
		FindByUserIdAndTaskIdByDay(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error)
		FindByUserIdAndTaskIdByWeek(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error)
	}

	customTaskRecordModel struct {
		*defaultTaskRecordModel
	}
)

func (m *defaultTaskRecordModel) FindByUserIdAndTaskId(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error) {
	var resp TaskRecord
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? AND task_id = ? ORDER BY id DESC LIMIT 1", taskRecordRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, taskId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTaskRecordModel) FindByUserId(ctx context.Context, userId int64, builder squirrel.SelectBuilder, orderBy string) ([]*TaskRecord, error) {
	builder = builder.Columns(taskRecordRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	// 获取今天的日期和本周的起始日期
	today := time.Now().Format("2006-01-02")
	weekStart := time.Now().AddDate(0, 0, -int(time.Now().Weekday())).Format("2006-01-02")

	query, values, err := builder.Where("user_id = ? AND (type = 1 OR (type = 2 AND DATE(create_time) = ?) OR (type = 3 AND DATE(create_time) >= ?))", userId, today, weekStart).ToSql()
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

func (m *defaultTaskRecordModel) FindByUserIdAndTaskIdByDay(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error) {
	var resp TaskRecord
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? AND task_id = ? AND DATE(create_time) = CURDATE() ORDER BY id DESC LIMIT 1", taskRecordRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, taskId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTaskRecordModel) FindByUserIdAndTaskIdByWeek(ctx context.Context, userId int64, taskId int64) (*TaskRecord, error) {
	var resp TaskRecord
	query := fmt.Sprintf("SELECT %s FROM %s WHERE user_id = ? AND task_id = ? AND WEEK(create_time) = WEEK(CURDATE()) ORDER BY id DESC LIMIT 1", taskRecordRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userId, taskId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
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
