// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
	"looklook/common/globalkey"
)

var (
	taskProgressFieldNames          = builder.RawFieldNames(&TaskProgress{})
	taskProgressRows                = strings.Join(taskProgressFieldNames, ",")
	taskProgressRowsExpectAutoSet   = strings.Join(stringx.Remove(taskProgressFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	taskProgressRowsWithPlaceHolder = strings.Join(stringx.Remove(taskProgressFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCheckinTaskProgressIdPrefix = "cache:checkin:taskProgress:id:"
)

type (
	taskProgressModel interface {
		Insert(ctx context.Context, data *TaskProgress) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *TaskProgress) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TaskProgress, error)
		Update(ctx context.Context, data *TaskProgress) error
		List(ctx context.Context, page, limit int64) ([]*TaskProgress, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *TaskProgress) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*TaskProgress, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TaskProgress, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TaskProgress, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*TaskProgress, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*TaskProgress, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultTaskProgressModel struct {
		sqlc.CachedConn
		table string
	}

	TaskProgress struct {
		Id                    int64 `db:"id"`
		UserId                int64 `db:"user_id"`
		IsParticipatedLottery int64 `db:"isParticipatedLottery"`
		IsInitiatedLottery    int64 `db:"isInitiatedLottery"`
		IsSubCheckin          int64 `db:"is_sub_checkin"`
	}
)

func newTaskProgressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTaskProgressModel {
	return &defaultTaskProgressModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`task_progress`",
	}
}

func (m *defaultTaskProgressModel) Delete(ctx context.Context, id int64) error {
	checkinTaskProgressIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, checkinTaskProgressIdKey)
	return err
}

func (m *defaultTaskProgressModel) FindOne(ctx context.Context, id int64) (*TaskProgress, error) {
	checkinTaskProgressIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressIdPrefix, id)
	var resp TaskProgress
	err := m.QueryRowCtx(ctx, &resp, checkinTaskProgressIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskProgressRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultTaskProgressModel) Insert(ctx context.Context, data *TaskProgress) (sql.Result, error) {
	checkinTaskProgressIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, taskProgressRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.IsParticipatedLottery, data.IsInitiatedLottery, data.IsSubCheckin)
	}, checkinTaskProgressIdKey)
	return ret, err
}

func (m *defaultTaskProgressModel) TransInsert(ctx context.Context, session sqlx.Session, data *TaskProgress) (sql.Result, error) {
	checkinTaskProgressIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, taskProgressRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.IsParticipatedLottery, data.IsInitiatedLottery, data.IsSubCheckin)
	}, checkinTaskProgressIdKey)
	return ret, err
}
func (m *defaultTaskProgressModel) Update(ctx context.Context, data *TaskProgress) error {
	checkinTaskProgressIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskProgressRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.IsParticipatedLottery, data.IsInitiatedLottery, data.IsSubCheckin, data.Id)
	}, checkinTaskProgressIdKey)
	return err
}

func (m *defaultTaskProgressModel) TransUpdate(ctx context.Context, session sqlx.Session, data *TaskProgress) error {
	checkinTaskProgressIdKey := fmt.Sprintf("%s%v", cacheCheckinTaskProgressIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, taskProgressRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.IsParticipatedLottery, data.IsInitiatedLottery, data.IsSubCheckin, data.Id)
	}, checkinTaskProgressIdKey)
	return err
}

func (m *defaultTaskProgressModel) List(ctx context.Context, page, limit int64) ([]*TaskProgress, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", taskProgressRows, m.table)
	var resp []*TaskProgress
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultTaskProgressModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultTaskProgressModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindSum Least One Field"), "FindSum Least One Field")
	}

	builder = builder.Columns("IFNULL(SUM(" + field + "),0)")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultTaskProgressModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

	if len(field) == 0 {
		return 0, errors.Wrapf(errors.New("FindCount Least One Field"), "FindCount Least One Field")
	}

	builder = builder.Columns("COUNT(" + field + ")")

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultTaskProgressModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*TaskProgress, error) {

	builder = builder.Columns(taskProgressRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TaskProgress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTaskProgressModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TaskProgress, error) {

	builder = builder.Columns(taskProgressRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TaskProgress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTaskProgressModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*TaskProgress, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(taskProgressRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, total, err
	}

	var resp []*TaskProgress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultTaskProgressModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*TaskProgress, error) {

	builder = builder.Columns(taskProgressRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TaskProgress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTaskProgressModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*TaskProgress, error) {

	builder = builder.Columns(taskProgressRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*TaskProgress
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTaskProgressModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultTaskProgressModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheCheckinTaskProgressIdPrefix, primary)
}

func (m *defaultTaskProgressModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", taskProgressRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTaskProgressModel) tableName() string {
	return m.table
}
