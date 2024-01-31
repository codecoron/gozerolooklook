package model

import (
	"context"
	"fmt"
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

	//自定义方法
	VoteConfigDiyModel interface {
		voteConfigModel
		QueryRows(ctx context.Context, whereString string) ([]*VoteConfig, error)
	}
)

// NewVoteConfigModel returns a model for the database table.
func NewVoteConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) VoteConfigModel {
	return &customVoteConfigModel{
		defaultVoteConfigModel: newVoteConfigModel(conn, c, opts...),
	}
}

func (c *customVoteConfigModel) QueryRows(ctx context.Context, whereString string) ([]*VoteConfig, error) {
	var resp []*VoteConfig
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", c.table, whereString)
	fmt.Println(query)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
