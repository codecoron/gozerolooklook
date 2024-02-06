package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSponsorModel = (*customUserSponsorModel)(nil)

type (
	// UserSponsorModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSponsorModel.
	UserSponsorModel interface {
		userSponsorModel
		//FindBySponsorId(ctx context.Context, SponsorId int64) ([]*UserSponsor, error)
	}

	customUserSponsorModel struct {
		*defaultUserSponsorModel
	}
)

//func (c customUserSponsorModel) FindBySponsorId(ctx context.Context, SponsorId int64) ([]*UserSponsor, error) {
//	//TODO implement me
//	var resp []*UserSponsor
//	query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", SponsorId)
//	panic("implement me")
//}

// NewUserSponsorModel returns a model for the database table.
func NewUserSponsorModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserSponsorModel {
	return &customUserSponsorModel{
		defaultUserSponsorModel: newUserSponsorModel(conn, c, opts...),
	}
}
