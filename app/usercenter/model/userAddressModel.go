package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAddressModel = (*customUserAddressModel)(nil)

type (
	// UserAddressModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAddressModel.
	UserAddressModel interface {
		userAddressModel

		List(ctx context.Context, page, limit int64) ([]*UserAddress, error)
		TransInsert(ctx context.Context, session sqlx.Session, data *UserAddress) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
	}

	customUserAddressModel struct {
		*defaultUserAddressModel
	}
)

// NewUserAddressModel returns a model for the database table.
func NewUserAddressModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserAddressModel {
	return &customUserAddressModel{
		defaultUserAddressModel: newUserAddressModel(conn, c, opts...),
	}
}

func (c *customUserAddressModel) List(ctx context.Context, page, limit int64) ([]*UserAddress, error) {
	query := fmt.Sprintf("select %s from %s limit ?,?", userAddressRows, c.table)
	var resp []*UserAddress
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, (page-1)*limit, limit)
	return resp, err
}

func (m *customUserAddressModel) TransInsert(ctx context.Context, session sqlx.Session, data *UserAddress) (sql.Result, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserAddressIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, userAddressRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.UserId, data.ContactName, data.ContactMobile, data.District, data.Detail, data.Postcode, data.IsDefault)
	}, lotteryLotteryIdKey)
	return ret, err
}

func (m *customUserAddressModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}
