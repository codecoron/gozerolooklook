package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		SetAdmin(ctx context.Context, uid int64) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (d *defaultUserModel) SetAdmin(ctx context.Context, uid int64) error {
	usercentUserIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserIdPrefix, uid)
	_, err := d.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set is_admin = CASE WHEN is_admin = 0 THEN 1 ELSE 0 END where `id` = ?", d.table)
		return conn.ExecCtx(ctx, query, uid)
	}, usercentUserIdKey)
	return err
}
