package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSponsorModel = (*customUserSponsorModel)(nil)

type (
	// UserSponsorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSponsorModel.
	UserSponsorModel interface {
		userSponsorModel
		FindPageByUserId(ctx context.Context, userId int64, offset int64, limit int64) ([]*UserSponsor, error)
	}

	customUserSponsorModel struct {
		*defaultUserSponsorModel
	}
)

// NewUserSponsorModel returns a model for the database table.
func NewUserSponsorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserSponsorModel {
	return &customUserSponsorModel{
		defaultUserSponsorModel: newUserSponsorModel(conn, c, opts...),
	}
}

func (m *defaultUserSponsorModel) FindPageByUserId(ctx context.Context, userId int64, offset int64, limit int64) ([]*UserSponsor, error) {
	var resp []*UserSponsor
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ? limit ?,?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userId, (offset-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
