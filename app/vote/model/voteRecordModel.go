package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ VoteRecordModel = (*customVoteRecordModel)(nil)

type (
	// VoteRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVoteRecordModel.
	VoteRecordModel interface {
		voteRecordModel
	}

	customVoteRecordModel struct {
		*defaultVoteRecordModel
	}
)

// NewVoteRecordModel returns a model for the database table.
func NewVoteRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VoteRecordModel {
	return &customVoteRecordModel{
		defaultVoteRecordModel: newVoteRecordModel(conn, c, opts...),
	}
}
