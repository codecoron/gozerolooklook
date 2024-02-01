package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
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

	//自定义方法
	VoteConfigDiyModel interface {
		voteConfigModel
		QueryRows(ctx context.Context, selectString string, whereString string) ([]*VoteConfig, error)
		QueryRow(ctx context.Context, selectString string, whereString string) (*VoteConfig, error)
	}
)

// NewVoteConfigModel returns a model for the database table.
func NewVoteConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VoteConfigModel {
	return &customVoteConfigModel{
		defaultVoteConfigModel: newVoteConfigModel(conn, c, opts...),
	}
}

func (c *customVoteConfigModel) QueryRows(ctx context.Context, selectString string, whereString string) ([]*VoteConfig, error) {
	if selectString == "" {
		selectString = "*"
	}
	var resp []*VoteConfig
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", selectString, c.table, whereString)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *customVoteConfigModel) QueryRow(ctx context.Context, selectString string, whereString string) (*VoteConfig, error) {
	if selectString == "" {
		selectString = "*"
	}
	var resp *VoteConfig
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", selectString, c.table, whereString)
	err := c.QueryRowNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
