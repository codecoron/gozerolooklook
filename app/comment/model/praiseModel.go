package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ PraiseModel = (*customPraiseModel)(nil)

type (
	// PraiseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPraiseModel.
	PraiseModel interface {
		praiseModel
		PraiseList(ctx context.Context, page, limit, lastId int64) ([]*Praise, error)
	}

	customPraiseModel struct {
		*defaultPraiseModel
	}
)

// NewPraiseModel returns a model for the database table.
func NewPraiseModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PraiseModel {
	return &customPraiseModel{
		defaultPraiseModel: newPraiseModel(conn, c, opts...),
	}
}

func (c *customPraiseModel) PraiseList(ctx context.Context, page, limit, lastId int64) ([]*Praise, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where id > ? limit ?,?", praiseRows, c.table)
	var resp []*Praise
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lastId, (page-1)*limit, limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, lastId:%v, (page-1)*limit:%v, limit:%v, error: %v", &resp, query, lastId, (page-1)*limit, limit, err)
	}
	return resp, nil
}
