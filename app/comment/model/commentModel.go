package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/globalkey"
	"looklook/common/xerr"
	"time"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		CommentList(ctx context.Context, page, limit, lastId, sort int64) ([]*Comment, error)
		UpdatePraiseNum(ctx context.Context, id, num int64) (int64, error)
		DeleteSoft(ctx context.Context, data *Comment) error
		GetCommentLastId() (int64, error)
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

func (c *customCommentModel) CommentList(ctx context.Context, page, limit, lastId, sort int64) ([]*Comment, error) {
	var query string
	if sort == 1 {
		// 按照点赞数倒序排序
		query = fmt.Sprintf("select %s from %s where id < ? order by praise_count desc limit ?,?", commentRows, c.table)
	} else {
		// 按照id倒序排序
		query = fmt.Sprintf("select %s from %s where id < ? order by id desc limit ?,?", commentRows, c.table)
	}
	var resp []*Comment
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

func (m *defaultCommentModel) DeleteSoft(ctx context.Context, data *Comment) error {
	data.DelState = globalkey.DelStateYes
	data.DeleteTime.Time = time.Now()
	data.DeleteTime.Valid = true
	if err := m.Update(ctx, data); err != nil {
		return errors.Wrapf(errors.New("delete soft failed "), "CommentModel delete err : %+v", err)
	}
	return nil
}

func (c *customCommentModel) GetCommentLastId() (int64, error) {
	var id int64
	query := fmt.Sprintf("select id from %s order by id desc limit 1", c.table)
	err := c.QueryRowNoCache(&id, query)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowNoCache, id:%v, query:%v, error: %v", id, query, err)
	}
	return id, nil
}
