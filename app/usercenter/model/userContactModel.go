package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserContactModel = (*customUserContactModel)(nil)

type (
	// UserContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserContactModel.
	UserContactModel interface {
		userContactModel
		FindPageByUserId(ctx context.Context, userId int64, offset int64, limit int64) ([]*UserContact, error)
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
