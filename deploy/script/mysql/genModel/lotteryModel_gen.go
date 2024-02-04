// Code generated by goctl. DO NOT EDIT.

package genModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

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
	lotteryFieldNames          = builder.RawFieldNames(&Lottery{})
	lotteryRows                = strings.Join(lotteryFieldNames, ",")
	lotteryRowsExpectAutoSet   = strings.Join(stringx.Remove(lotteryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	lotteryRowsWithPlaceHolder = strings.Join(stringx.Remove(lotteryFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLotteryLotteryIdPrefix = "cache:lottery:lottery:id:"
)

type (
	lotteryModel interface {
		Insert(ctx context.Context, data *Lottery) (sql.Result, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *Lottery) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Lottery, error)
		Update(ctx context.Context, data *Lottery) error
		List(ctx context.Context, page, limit int64) ([]*Lottery, error)
		TransUpdate(ctx context.Context, session sqlx.Session, data *Lottery) error
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		SelectBuilder() squirrel.SelectBuilder
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder, field string) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder, field string) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*Lottery, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Lottery, error)
		FindPageListByPageWithTotal(ctx context.Context, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Lottery, int64, error)
		FindPageListByIdDESC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Lottery, error)
		FindPageListByIdASC(ctx context.Context, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Lottery, error)
		Delete(ctx context.Context, id int64) error
	}

	defaultLotteryModel struct {
		sqlc.CachedConn
		table string
	}

	Lottery struct {
		Id            int64        `db:"id"`
		UserId        int64        `db:"user_id"`        // 发起抽奖用户ID
		Name          string       `db:"name"`           // 默认取一等奖名称
		Thumb         string       `db:"thumb"`          // 默认取一等经配图
		PublishTime   sql.NullTime `db:"publish_time"`   // 发布抽奖时间
		JoinNumber    int64        `db:"join_number"`    // 自动开奖人数
		Introduce     string       `db:"introduce"`      // 抽奖说明
		AwardDeadline time.Time    `db:"award_deadline"` // 领奖截止时间
		IsSelected    int64        `db:"is_selected"`    // 是否精选 1是 0否
		AnnounceType  int64        `db:"announce_type"`  // 开奖设置：1按时间开奖 2按人数开奖 3即抽即中
		AnnounceTime  sql.NullTime `db:"announce_time"`  // 开奖时间
		IsAnnounced   int64        `db:"is_announced"`   // 是否开奖：0未开奖；1已经开奖
		CreateTime    time.Time    `db:"create_time"`
		UpdateTime    time.Time    `db:"update_time"`
		DeleteTime    sql.NullTime `db:"delete_time"` // 删除时间
		SponsorId     int64        `db:"sponsor_id"`  // 发起抽奖赞助商ID
	}
)

func newLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultLotteryModel {
	return &defaultLotteryModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`lottery`",
	}
}

func (m *defaultLotteryModel) Delete(ctx context.Context, id int64) error {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, lotteryLotteryIdKey)
	return err
}

func (m *defaultLotteryModel) FindOne(ctx context.Context, id int64) (*Lottery, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, id)
	var resp Lottery
	err := m.QueryRowCtx(ctx, &resp, lotteryLotteryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", lotteryRows, m.table)
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

func (m *defaultLotteryModel) Insert(ctx context.Context, data *Lottery) (sql.Result, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, lotteryRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.Name, data.Thumb, data.PublishTime, data.JoinNumber, data.Introduce, data.AwardDeadline, data.IsSelected, data.AnnounceType, data.AnnounceTime, data.IsAnnounced, data.DeleteTime, data.SponsorId)
	}, lotteryLotteryIdKey)
	return ret, err
}

func (m *defaultLotteryModel) TransInsert(ctx context.Context, session sqlx.Session, data *Lottery) (sql.Result, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, lotteryRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.Name, data.Thumb, data.PublishTime, data.JoinNumber, data.Introduce, data.AwardDeadline, data.IsSelected, data.AnnounceType, data.AnnounceTime, data.IsAnnounced, data.DeleteTime, data.SponsorId)
	}, lotteryLotteryIdKey)
	return ret, err
}
func (m *defaultLotteryModel) Update(ctx context.Context, data *Lottery) error {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, lotteryRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.Name, data.Thumb, data.PublishTime, data.JoinNumber, data.Introduce, data.AwardDeadline, data.IsSelected, data.AnnounceType, data.AnnounceTime, data.IsAnnounced, data.DeleteTime, data.SponsorId, data.Id)
	}, lotteryLotteryIdKey)
	return err
}

func (m *defaultLotteryModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Lottery) error {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, lotteryRowsWithPlaceHolder)
		return session.ExecCtx(ctx, query, data.UserId, data.Name, data.Thumb, data.PublishTime, data.JoinNumber, data.Introduce, data.AwardDeadline, data.IsSelected, data.AnnounceType, data.AnnounceTime, data.IsAnnounced, data.DeleteTime, data.SponsorId, data.Id)
	}, lotteryLotteryIdKey)
	return err
}

func (m *defaultLotteryModel) List(ctx context.Context, page, limit int64) ([]*Lottery, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", lotteryRows, m.table)
	var resp []*Lottery
	//err := m.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *defaultLotteryModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultLotteryModel) FindSum(ctx context.Context, builder squirrel.SelectBuilder, field string) (float64, error) {

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

func (m *defaultLotteryModel) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64, error) {

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

func (m *defaultLotteryModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Lottery, error) {

	builder = builder.Columns(lotteryRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Lottery
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultLotteryModel) FindPageListByPage(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Lottery, error) {

	builder = builder.Columns(lotteryRows)

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

	var resp []*Lottery
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultLotteryModel) FindPageListByPageWithTotal(ctx context.Context, builder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*Lottery, int64, error) {

	total, err := m.FindCount(ctx, builder, "id")
	if err != nil {
		return nil, 0, err
	}

	builder = builder.Columns(lotteryRows)

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

	var resp []*Lottery
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, total, nil
	default:
		return nil, total, err
	}
}

func (m *defaultLotteryModel) FindPageListByIdDESC(ctx context.Context, builder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*Lottery, error) {

	builder = builder.Columns(lotteryRows)

	if preMinId > 0 {
		builder = builder.Where(" id < ? ", preMinId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Lottery
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultLotteryModel) FindPageListByIdASC(ctx context.Context, builder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*Lottery, error) {

	builder = builder.Columns(lotteryRows)

	if preMaxId > 0 {
		builder = builder.Where(" id > ? ", preMaxId)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Lottery
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultLotteryModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}

func (m *defaultLotteryModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, primary)
}

func (m *defaultLotteryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", lotteryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLotteryModel) tableName() string {
	return m.table
}
