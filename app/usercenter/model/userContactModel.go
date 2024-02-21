package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strconv"
	"strings"
)

var _ UserContactModel = (*customUserContactModel)(nil)

type (
	// UserContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserContactModel.
	UserContactModel interface {
		userContactModel
		FindPageByUserId(ctx context.Context, userId int64, offset int64, limit int64) ([]*UserContact, error)
		DeleteBatch(ctx context.Context, id []int64) error
	}

	customUserContactModel struct {
		*defaultUserContactModel
	}
)

// NewUserContactModel returns a model for the database table.
func NewUserContactModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserContactModel {
	return &customUserContactModel{
		defaultUserContactModel: newUserContactModel(conn, c, opts...),
	}
}

func (m *defaultUserContactModel) FindPageByUserId(ctx context.Context, userId int64, offset int64, limit int64) ([]*UserContact, error) {
	var resp []*UserContact
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ? limit ?,?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId, (offset-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultUserContactModel) DeleteBatch(ctx context.Context, ids []int64) error {
	////todo 优化这里的逻辑
	stringSlice := make([]string, len(ids))
	for i, num := range ids {
		stringSlice[i] = strconv.FormatInt(num, 10)
	}
	//todo 特殊处理缓存
	looklookUsercenterUserContactIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserContactIdPrefix, ids)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		//query := fmt.Sprintf("delete from %s where `id` in ?", m.table)
		query := fmt.Sprintf("SELECT * FROM table WHERE id IN (%s)", m.table)
		return conn.ExecCtx(ctx, query, strings.Join(stringSlice, ","))
	}, looklookUsercenterUserContactIdKey)
	return err
}
