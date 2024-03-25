package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"
	"strconv"
	"strings"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		SetAdmin(ctx context.Context, uid int64) error
		FindUserInfoByUserIds(ctx context.Context, userIds []int64) ([]*pb.UserInfoForComment, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (d *defaultUserModel) SetAdmin(ctx context.Context, uid int64) error {
	usercentUserIdKey := fmt.Sprintf("%s%v", cacheLooklookUsercenterUserIdPrefix, uid)
	_, err := d.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set is_admin = CASE WHEN is_admin = 0 THEN 1 ELSE 0 END where `id` = ?", d.table)
		return conn.ExecCtx(ctx, query, uid)
	}, usercentUserIdKey)
	return err
}

func (d *defaultUserModel) FindUserInfoByUserIds(ctx context.Context, userIds []int64) ([]*pb.UserInfoForComment, error) {

	//query := fmt.Sprintf("SELECT * FROM %s WHERE id = ? ", m.table)
	idStrings := make([]string, 0, len(userIds))
	for _, id := range userIds {
		idStrings = append(idStrings, strconv.FormatInt(id, 10))
	}
	commaSeparated := strings.Join(idStrings, ",")

	var resp []*pb.UserInfoForComment
	query := fmt.Sprintf("select id,nickname,avatar from %s where id in (%s)", d.table, commaSeparated)
	logx.Debug("rpc为评论查询用户信息", query)
	err := d.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "ExecCtx, query:%v,  commaSeparated:%v, error: %v", query, commaSeparated, err)
	}

	return resp, nil
}
