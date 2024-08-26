package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
)

var _ PrizeModel = (*customPrizeModel)(nil)
var (
	cacheLotteryPrizeIdLotteryIdPrefix = "cache:lottery:prize:id:lotteryId:"
)

type (
	// PrizeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPrizeModel.
	PrizeModel interface {
		prizeModel
		//自定义的方法写道这里，避免覆盖
		TransInsert(ctx context.Context, session sqlx.Session, data *Prize) (sql.Result, error)
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		FindByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error)
		FindOneByIdLotteryId(ctx context.Context, id int64, lotteryId int64) (*Prize, error)
		FindPageByLotteryId(ctx context.Context, lotteryId int64, offset int64, limit int64) ([]*Prize, error)
		GetPrizeInfoByPrizeIds(ctx context.Context, prizeIds []int64) ([]*Prize, error)
		FindAllByLotteryIds(ctx context.Context, lotteryIds []int64) ([]*Prize, error)
	}

	customPrizeModel struct {
		*defaultPrizeModel
	}
)

// NewPrizeModel returns a model for the database table.
func NewPrizeModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PrizeModel {
	return &customPrizeModel{
		defaultPrizeModel: newPrizeModel(conn, c, opts...),
	}
}

func (m *defaultPrizeModel) TransInsert(ctx context.Context, session sqlx.Session, data *Prize) (sql.Result, error) {
	lotteryPrizeIdKey := fmt.Sprintf("%s%v", cacheLotteryPrizeIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, prizeRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.LotteryId, data.Type, data.Name, data.Level, data.Thumb, data.Count, data.GrantType)
	}, lotteryPrizeIdKey)
	return ret, err
}

func (m *defaultPrizeModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultPrizeModel) FindByLotteryId(ctx context.Context, lotteryId int64) ([]*Prize, error) {
	var resp []*Prize
	query := fmt.Sprintf("SELECT * FROM %s WHERE lottery_id = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_PRIZES_BYLOTTERYID_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, lotteryId:%v, error: %v", &resp, query, lotteryId, err)
	}
	return resp, nil
}

func (m *defaultPrizeModel) FindOneByIdLotteryId(ctx context.Context, id int64, lotteryId int64) (*Prize, error) {
	lotteryPrizeIdLotteryIdKey := fmt.Sprintf("%s%v:%v", cacheLotteryPrizeIdLotteryIdPrefix, id, lotteryId)
	var resp Prize
	err := m.QueryRowIndexCtx(ctx, &resp, lotteryPrizeIdLotteryIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `id` = ? and `lottery_id` = ? limit 1", prizeRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, id, lotteryId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultPrizeModel) FindPageByLotteryId(ctx context.Context, lotteryId int64, offset int64, limit int64) ([]*Prize, error) {
	var resp []*Prize
	query := fmt.Sprintf("SELECT * FROM %s WHERE lottery_id = ? limit ?,?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, lotteryId, (offset-1)*limit, limit)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultPrizeModel) GetPrizeInfoByPrizeIds(ctx context.Context, prizeIds []int64) ([]*Prize, error) {
	var resp []*Prize
	// 这里传int64类型的切片，需要将切片转换成字符串，然后在sql语句中使用in关键字
	prizeIdsStr := ""
	for i, v := range prizeIds {
		if i == 0 {
			prizeIdsStr = fmt.Sprintf("%d", v)
		} else {
			prizeIdsStr = fmt.Sprintf("%s,%d", prizeIdsStr, v)
		}
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE id in (%s)", m.table, prizeIdsStr)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultPrizeModel) FindAllByLotteryIds(ctx context.Context, lotteryIds []int64) ([]*Prize, error) {
	var resp []*Prize
	// 这里传int64类型的切片，需要将切片转换成字符串，然后在sql语句中使用in关键字
	lotteryIdsStr := ""
	for i, v := range lotteryIds {
		if i == 0 {
			lotteryIdsStr = fmt.Sprintf("%d", v)
		} else {
			lotteryIdsStr = fmt.Sprintf("%s,%d", lotteryIdsStr, v)
		}
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE lottery_id in (%s)", m.table, lotteryIdsStr)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
