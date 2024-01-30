package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LotteryParticipationModel = (*customLotteryParticipationModel)(nil)

type (
	// LotteryParticipationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryParticipationModel.
	LotteryParticipationModel interface {
		lotteryParticipationModel
		GetParticipationUserIdsByLotteryId(ctx context.Context, LotteryId int64) ([]int64, error)
		UpdateWinners(ctx context.Context, LotteryId, UserId, PrizeId int64) error
		GetParticipatorsCountByLotteryId(ctx context.Context, LotteryId int64) (int64, error)
	}

	customLotteryParticipationModel struct {
		*defaultLotteryParticipationModel
	}
)

// NewLotteryParticipationModel returns a model for the database table.
func NewLotteryParticipationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryParticipationModel {
	return &customLotteryParticipationModel{
		defaultLotteryParticipationModel: newLotteryParticipationModel(conn, c, opts...),
	}
}

func (m *defaultLotteryParticipationModel) UpdateWinners(ctx context.Context, LotteryId, UserId, PrizeId int64) error {
	query := fmt.Sprintf("update %s set is_won = 1, prize_id = ? where `lottery_id` = ? and `user_id` = ?", m.table)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		res, err := conn.ExecCtx(ctx, query, PrizeId, LotteryId, UserId)
		if err != nil {
			return nil, err
		}
		return res, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultLotteryParticipationModel) GetParticipationUserIdsByLotteryId(ctx context.Context, LotteryId int64) ([]int64, error) {
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE lottery_id = ?", m.table)
	var resp []int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, LotteryId)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (m *defaultLotteryParticipationModel) GetParticipatorsCountByLotteryId(ctx context.Context, LotteryId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE lottery_id = ?", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, LotteryId)
	if err != nil {
		return 0, err
	}
	return resp, nil
}
