package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"
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
		CheckIsWonByUserIdAndLotteryId(ctx context.Context, LotteryId, UserId int64) (int64, error)
		GetWonListByUserId(ctx context.Context, UserId, Page, Size, LastId int64) ([]*LotteryParticipation, error)
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
		return errors.Wrapf(xerr.NewErrCode(xerr.UPDATE_WINNER_ERROR), "UpdateWinners, PrizeId:%v, LotteryId:%v, UserId:%v, error: %v", PrizeId, LotteryId, UserId, err)
	}
	return nil
}

func (m *defaultLotteryParticipationModel) GetParticipationUserIdsByLotteryId(ctx context.Context, LotteryId int64) ([]int64, error) {
	query := fmt.Sprintf("SELECT user_id FROM %s WHERE lottery_id = ?", m.table)
	var resp []int64
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, LotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_PARTICIPATION_USERIDS_BYLOTTERYID_ERROR), "GetParticipationUserIdsByLotteryId,LotteryId:%v, error: %v", LotteryId, err)
	}
	return resp, nil
}

func (m *defaultLotteryParticipationModel) GetParticipatorsCountByLotteryId(ctx context.Context, LotteryId int64) (int64, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE lottery_id = ?", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, LotteryId)
	if err != nil {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.GET_PARTICIPATORS_COUNT_BYLOTTERYID_ERROR), "GetParticipatorsCountByLotteryId, LotteryId:%v, error: %v", LotteryId, err)
	}
	return resp, nil
}

func (m *defaultLotteryParticipationModel) CheckIsWonByUserIdAndLotteryId(ctx context.Context, LotteryId, UserId int64) (int64, error) {
	query := fmt.Sprintf("SELECT is_won FROM %s WHERE lottery_id = ? AND user_id = ?", m.table)
	var resp int64
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, LotteryId, UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.CHECK_ISWON_BYUSERID_ANDLOTTERYID_ERROR), "CheckIsWonByUserIdAndLotteryId, LotteryId:%v, UserId:%v, error: %v", LotteryId, UserId, err)
	}
	return resp, nil
}

func (m *defaultLotteryParticipationModel) GetWonListByUserId(ctx context.Context, UserId, Page, Size, LastId int64) ([]*LotteryParticipation, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = ? AND is_won = 1 AND id > ? LIMIT ?, ?", m.table)
	var resp []*LotteryParticipation
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, UserId, LastId, (Page-1)*Size, Size)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_WONLIST_BYUSERID_ERROR), "GetWonListByUserId, UserId:%v, Page:%v, Size:%v, error: %v", UserId, Page, Size, err)
	}
	return resp, nil
}
