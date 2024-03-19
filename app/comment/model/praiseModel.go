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
		GetLikeCountByCommentIds(ctx context.Context, commentIds []int64) (map[int64]int64, error)
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
	// 这里传int64类型的切片，需要将切片转换成字符串，然后在sql语句中使用in关键字
	commentIdsStr := ""
	for i, v := range commentIds {
		if i == 0 {
			commentIdsStr = fmt.Sprintf("%d", v)
		} else {
			commentIdsStr = fmt.Sprintf("%s,%d", commentIdsStr, v)
		}
	}
	query := fmt.Sprintf("select comment_id from %s where user_id = ? and comment_id in (%s)", c.table, commentIdsStr)

	err := c.QueryRowsNoCacheCtx(ctx, &ids, query, userId)
	if err != nil && err != sqlx.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &ids:%v, query:%v, commentIds:%v, userId:%v, error: %v", &ids, query, commentIds, userId, err)
	}

	return ids, nil
}

// PraiseCount 新建一个结构体，接收查询结果
type PraiseCount struct {
	CommentId int64
	Count     int64
}

func (c *customPraiseModel) GetLikeCountByCommentIds(ctx context.Context, commentIds []int64) (map[int64]int64, error) {
	// 查询评论的点赞数
	likeCount := make(map[int64]int64)
	// 这里传int64类型的切片，需要将切片转换成字符串，然后在sql语句中使用in关键字
	commentIdsStr := ""
	for i, v := range commentIds {
		if i == 0 {
			commentIdsStr = fmt.Sprintf("%d", v)
		} else {
			commentIdsStr = fmt.Sprintf("%s,%d", commentIdsStr, v)
		}
	}
	query := fmt.Sprintf("SELECT comment_id, count(*) as count FROM %s WHERE comment_id in (%s) GROUP BY comment_id", c.table, commentIdsStr)
	var list []*PraiseCount
	err := c.QueryRowsNoCacheCtx(ctx, &list, query)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "QueryRowsNoCacheCtx, &list:%v, query:%v, error: %v", &list, query, err)
	}
	for _, v := range list {
		likeCount[v.CommentId] = v.Count
	}
	return likeCount, nil
}
