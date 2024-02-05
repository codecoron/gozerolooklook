package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	_                                    TaskProgressModel = (*customTaskProgressModel)(nil)
	cacheCheckinTaskProgressUserIdPrefix                   = "cache:checkin:taskProgress:userId:"
)

type (
	// TaskProgressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTaskProgressModel.
	TaskProgressModel interface {
		taskProgressModel
		FindOneByUserId(ctx context.Context, userId int64) (*TaskProgress, error)
		InsertByUserId(ctx context.Context, data *TaskProgress) (sql.Result, error)
		TransUpdateByUserId(ctx context.Context, session sqlx.Session, data *TaskProgress) error
		UpdateByUserId(ctx context.Context, data *TaskProgress) error
		FindAllSubId(ctx context.Context) ([]int64, error)
	}

	customTaskProgressModel struct {
		*defaultTaskProgressModel
	}
)

func (m *defaultTaskProgressModel) FindOneByUserId(ctx context.Context, userId int64) (*TaskProgress, error) {
	checkinTaskProgressUserIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressUserIdPrefix, userId)
	var resp TaskProgress
	err := m.QueryRowCtx(ctx, &resp, checkinTaskProgressUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", taskProgressRows, m.table)
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

func (m *defaultTaskProgressModel) InsertByUserId(ctx context.Context, data *TaskProgress) (sql.Result, error) {
	checkinTaskProgressUserIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, taskProgressRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.IsParticipatedLottery, data.IsInitiatedLottery, data.IsSubCheckin)
	}, checkinTaskProgressUserIdKey)
	return ret, err
}

func (m *defaultTaskProgressModel) TransUpdateByUserId(ctx context.Context, session sqlx.Session, data *TaskProgress) error {
	checkinTaskProgressUserIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressUserIdPrefix, data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskProgressRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.IsParticipatedLottery, data.IsInitiatedLottery, data.IsSubCheckin, data.Id)
	}, checkinTaskProgressUserIdKey)
	return err
}

func (m *defaultTaskProgressModel) UpdateByUserId(ctx context.Context, data *TaskProgress) error {
	checkinTaskProgressUserIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressUserIdPrefix, data.UserId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskProgressRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.IsParticipatedLottery, data.IsInitiatedLottery, data.IsSubCheckin, data.Id)
	}, checkinTaskProgressUserIdKey)
	return err
}

func (m *defaultTaskProgressModel) FindAllSubId(ctx context.Context) ([]int64, error) {
	builder := squirrel.Select("user_id").
		From(m.table).
		Where("is_sub_checkin = 1")

	query, values, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var userIDs []int64
	err = m.QueryRowsNoCacheCtx(ctx, &userIDs, query, values...)
	switch err {
	case nil:
		return userIDs, nil
	default:
		return nil, err
	}
}

// NewTaskProgressModel returns a model for the database table.
func NewTaskProgressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) TaskProgressModel {
	return &customTaskProgressModel{
		defaultTaskProgressModel: newTaskProgressModel(conn, c, opts...),
	}
}
