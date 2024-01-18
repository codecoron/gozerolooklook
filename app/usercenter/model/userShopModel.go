package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserShopModel = (*customUserShopModel)(nil)

type (
	// UserShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserShopModel.
	UserShopModel interface {
		userShopModel
	}

	customUserShopModel struct {
		*defaultUserShopModel
	}
)

// NewUserShopModel returns a model for the database table.
func NewUserShopModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserShopModel {
	return &customUserShopModel{
		defaultUserShopModel: newUserShopModel(conn, c, opts...),
	}
}
