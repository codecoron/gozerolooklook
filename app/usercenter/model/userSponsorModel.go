package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSponsorModel = (*customUserSponsorModel)(nil)

type (
	// UserSponsorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSponsorModel.
	UserSponsorModel interface {
		userSponsorModel
	}

	customUserSponsorModel struct {
		*defaultUserSponsorModel
	}
)

// NewUserSponsorModel returns a model for the database table.
func NewUserSponsorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserSponsorModel {
	return &customUserSponsorModel{
		defaultUserSponsorModel: newUserSponsorModel(conn, c, opts...),
	}
}
