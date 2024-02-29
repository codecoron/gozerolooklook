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
		IsPraise(ctx context.Context, commentId, userId int64) (int64, error)
		IsPraiseThisWeek(ctx context.Context, userId int64) (bool, error)
		IsPraiseList(ctx context.Context, commentIds []int64, userId int64) ([]int64, error)
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
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lastId, (page-1)*limit, limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, lastId:%v, (page-1)*limit:%v, limit:%v, error: %v", &resp, query, lastId, (page-1)*limit, limit, err)
	}
	return resp, nil
}

func (c *customPraiseModel) IsPraise(ctx context.Context, commentId, userId int64) (int64, error) {
	// 查询是否有点赞记录，有则返回点赞id，否则返回0
	var id int64
	query := fmt.Sprintf("select id from %s where comment_id = ? and user_id = ? limit 1", c.table)
	err := c.QueryRowNoCacheCtx(ctx, &id, query, commentId, userId)
	if err != nil && err != sqlx.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &id:%v, query:%v, commentId:%v, userId:%v, error: %v", &id, query, commentId, userId, err)
	}
	return id, nil
}

func (c *customPraiseModel) IsPraiseThisWeek(ctx context.Context, userId int64) (bool, error) {
	// 查询是否有最新的点赞记录，有则返回 true，否则返回 false
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE user_id = ? AND YEARWEEK(create_time) = YEARWEEK(CURDATE()) ORDER BY create_time DESC LIMIT 1)", c.table)
	var exists bool
	err := c.QueryRowNoCacheCtx(ctx, &exists, query, userId)
	if err != nil && err != sqlx.ErrNotFound {
		return false, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowNoCacheCtx, exists:%v, query:%v, userId:%v, error: %v", exists, query, userId, err)
	}
	return exists, nil
}

func (c *customPraiseModel) IsPraiseList(ctx context.Context, commentIds []int64, userId int64) ([]int64, error) {
	// 查询是否有点赞记录，有则返回点赞id，否则返回0
	var ids []int64
	query := fmt.Sprintf("select comment_id from %s where comment_id in (?) and user_id = ?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &ids, query, commentIds, userId)
	if err != nil && err != sqlx.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &ids:%v, query:%v, commentIds:%v, userId:%v, error: %v", &ids, query, commentIds, userId, err)
	}
	return ids, nil
}
