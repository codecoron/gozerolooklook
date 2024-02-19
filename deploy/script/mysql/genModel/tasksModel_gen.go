// Code generated by goctl. DO NOT EDIT.

package genModel

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
	tasksFieldNames          = builder.RawFieldNames(&Tasks{})
	tasksRows                = strings.Join(tasksFieldNames, ",")
	tasksRowsExpectAutoSet   = strings.Join(stringx.Remove(tasksFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	tasksRowsWithPlaceHolder = strings.Join(stringx.Remove(tasksFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheCheckinTasksIdPrefix = "cache:checkin:tasks:id:"
)

type (
	tasksModel interface {
		Insert(ctx context.Context, data *Tasks) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Tasks) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Tasks, error)
		Update(ctx context.Context, data *Tasks) error
		List(ctx context.Context, page, limit int64) ([]*Tasks, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *Tasks) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Tasks, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Tasks, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Tasks, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Tasks, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Tasks, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultTasksModel struct {
		sqlc.CachedConn
		table string
	}

	Tasks struct {
		Id       int64  `db:"id"`
		Type     int64  `db:"type"` // 1 for novice, 2 for daily, 3 for weekly
		Content  string `db:"content"`
		Integral int64  `db:"integral"` // Increased wish value after completion
		DelState int64  `db:"del_state"`
	}
)

func newTasksModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultTasksModel {
	return &defaultTasksModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`tasks`",
	}
}

func (m *defaultTasksModel) Delete(ctx context.Context, id int64) error {
	checkinTasksIdKey := fmt.Sprintf("%s%v", cacheCheckinTasksIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, checkinTasksIdKey)
	return err
}

func (m *defaultTasksModel) FindOne(ctx context.Context, id int64) (*Tasks, error) {
	checkinTasksIdKey := fmt.Sprintf("%s%v", cacheCheckinTasksIdPrefix, id)
	var resp Tasks
	err := m.QueryRowCtx(ctx, &resp, checkinTasksIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tasksRows, m.table)
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

func (m *defaultTasksModel) Insert(ctx context.Context, data *Tasks) (sql.Result, error) {
	checkinTasksIdKey := fmt.Sprintf("%s%v", cacheCheckinTasksIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, tasksRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Type, data.Content, data.Integral, data.DelState)
	}, checkinTasksIdKey)
	return ret, err
}

func (m *defaultTasksModel) TransInsert(ctx context.Context, session sqlx.Session, data *Tasks) (sql.Result, error) {
	checkinTasksIdKey := fmt.Sprintf("%s%v", cacheCheckinTasksIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, tasksRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Type, data.Content, data.Integral, data.DelState)
	}, checkinTasksIdKey)
	return ret, err
}
func (m *defaultTasksModel) Update(ctx context.Context, data *Tasks) error {
	checkinTasksIdKey := fmt.Sprintf("%s%v", cacheCheckinTasksIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tasksRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Type, data.Content, data.Integral, data.DelState, data.Id)
	}, checkinTasksIdKey)
	return err
}

func (m *defaultTasksModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Tasks) error {
	checkinTasksIdKey := fmt.Sprintf("%s%v", cacheCheckinTasksIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tasksRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.Type, data.Content, data.Integral, data.DelState, data.Id)
	}, checkinTasksIdKey)
	return err
}

func (m *defaultTasksModel) List(ctx context.Context, page, limit int64) ([]*Tasks, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", tasksRows, m.table)
	var resp []*Tasks
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultTasksModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultTasksModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultTasksModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultTasksModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Tasks, error) {

	builder = builder.Columns(tasksRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Tasks
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTasksModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Tasks, error) {

	builder = builder.Columns(tasksRows)

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

	var resp []*Tasks
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTasksModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Tasks, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(tasksRows)

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

	var resp []*Tasks
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultTasksModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Tasks, error) {

	builder = builder.Columns(tasksRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Tasks
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTasksModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Tasks, error) {

	builder = builder.Columns(tasksRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Tasks
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultTasksModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultTasksModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheCheckinTasksIdPrefix, primary)
}

func (m *defaultTasksModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tasksRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTasksModel) tableName() string {
	return m.table
}
