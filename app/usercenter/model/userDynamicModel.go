package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ UserDynamicModel = (*customUserDynamicModel)(nil)

type (
	// UserDynamicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserDynamicModel.
	UserDynamicModel interface {
		userDynamicModel
		// 自定义方法
		FindListByUserId(ctx context.Context, userId int64) ([]*UserDynamic, error)
	}

	customUserDynamicModel struct {
		*defaultUserDynamicModel
	}
)

func (c *customUserDynamicModel) FindListByUserId(ctx context.Context, userId int64) ([]*UserDynamic, error) {
	var resp []*UserDynamic
	query := fmt.Sprintf("select * from %s where user_id = %d order by create_time DESC", c.table, userId)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query)
	//for i, dynamic := range resp {
	//	resp[i].UpdateTime = time.Now().
	//}
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_TYPEIS2_AND_ISNOT_ANNOUNCE_LOTTERYS_ERROR), "QueryRowsNoCacheCtx,&resp:%v, query:%v, error: %v", &resp, query, err)
	}
	return resp, nil
}

// NewUserDynamicModel returns a model for the database table.
func NewUserDynamicModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserDynamicModel {
	return &customUserDynamicModel{
		defaultUserDynamicModel: newUserDynamicModel(conn, c, opts...),
	}
}
