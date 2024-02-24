package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		CommentList(ctx context.Context, page, limit, lastId int64) ([]*Comment, error)
		UpdatePraiseNum(ctx context.Context, id, num int64) (int64, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c, opts...),
	}
}

func (c *customCommentModel) CommentList(ctx context.Context, page, limit, lastId int64) ([]*Comment, error) {
	var query string
	query = fmt.Sprintf("select %s from %s where id > ? limit ?,?", commentRows, c.table)
	var resp []*Comment
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lastId, (page-1)*limit, limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, lastId:%v, (page-1)*limit:%v, limit:%v, error: %v", &resp, query, lastId, (page-1)*limit, limit, err)
	}
	return resp, nil
}

func (c *customCommentModel) UpdatePraiseNum(ctx context.Context, id, num int64) (int64, error) {
	query := fmt.Sprintf("update %s set praise_count = praise_count + ? where id = ?", c.table)
	res, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, num, id)
	})
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ExecCtx, query:%v, num:%v, id:%v, error: %v", query, num, id, err)
	}
	return res.RowsAffected()
}
