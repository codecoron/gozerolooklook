package genModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VoteConfigModel = (*customVoteConfigModel)(nil)

type (
	// VoteConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVoteConfigModel.
	VoteConfigModel interface {
		voteConfigModel
	}

	customVoteConfigModel struct {
		*defaultVoteConfigModel
	}
)

// NewVoteConfigModel returns a model for the database table.
func NewVoteConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VoteConfigModel {
	return &customVoteConfigModel{
		defaultVoteConfigModel: newVoteConfigModel(conn, c, opts...),
	}
}
