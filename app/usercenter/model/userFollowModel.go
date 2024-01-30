package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserFollowModel = (*customUserFollowModel)(nil)

type (
	// UserFollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserFollowModel.
	UserFollowModel interface {
		userFollowModel
	}

	customUserFollowModel struct {
		*defaultUserFollowModel
	}
)

// NewUserFollowModel returns a model for the database table.
func NewUserFollowModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserFollowModel {
	return &customUserFollowModel{
		defaultUserFollowModel: newUserFollowModel(conn, c, opts...),
	}
}
