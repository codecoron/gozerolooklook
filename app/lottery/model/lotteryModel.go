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
	"time"
)

var _ LotteryModel = (*customLotteryModel)(nil)

type (
	// LotteryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLotteryModel.
	LotteryModel interface {
		lotteryModel
		// 自定义方法
		UpdatePublishTime(ctx context.Context, data *Lottery) error
		LotteryList(ctx context.Context, page, limit, selected, lastId int64) ([]*Lottery, error)
		FindUserIdByLotteryId(ctx context.Context, lotteryId int64) (*int64, error)
		GetLotterysByLessThanCurrentTime(ctx context.Context, currentTime time.Time, announceType int64) ([]int64, error)
		UpdateLotteryStatus(ctx context.Context, lotteryID int64) error
		GetTypeIs2AndIsNotAnnounceLotterys(ctx context.Context, announceType int64) ([]*Lottery, error)
		GetLotteryIdByUserId(ctx context.Context, UserId int64) (*int64, error)
		GetTodayLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error)
		GetWeekLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error)
	}

	customLotteryModel struct {
		*defaultLotteryModel
	}
)

func (m *defaultLotteryModel) UpdatePublishTime(ctx context.Context, data *Lottery) error {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set publish_time = ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, data.PublishTime, data.Id)
	}, lotteryLotteryIdKey)
	return err
}

func (c *customLotteryModel) LotteryList(ctx context.Context, page, limit, selected, lastId int64) ([]*Lottery, error) {
	var query string
	if selected != 0 {
		query = fmt.Sprintf("select %s from %s where id > ? and is_selected = 1 limit ?,?", lotteryRows, c.table)
	} else {
		query = fmt.Sprintf("select %s from %s where id > ? limit ?,?", lotteryRows, c.table)
	}
	var resp []*Lottery
	//err := c.conn.QueryRowsCtx(ctx, &resp, query, (page-1)*limit, limit)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, lastId, (page-1)*limit, limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_USERID_BYLOTTERYID_ERROR), "QueryRowsNoCacheCtx, &resp:%v, query:%v, lastId:%v, (page-1)*limit:%v, limit:%v, error: %v", &resp, query, lastId, (page-1)*limit, limit, err)
	}
	return resp, nil
}

func (c *customLotteryModel) FindUserIdByLotteryId(ctx context.Context, lotteryId int64) (*int64, error) {
	lotteryLotteryIdKey := fmt.Sprintf("%s%v", cacheLotteryLotteryIdPrefix, lotteryId)
	var resp int64
	err := c.QueryRowCtx(ctx, &resp, lotteryLotteryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select user_id from %s where id = ?", c.table)
		return conn.QueryRowCtx(ctx, v, query, lotteryId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_USERID_NOTFOUND), "FindUserIdByLotteryId, lotteryId:%v, error: %v", lotteryId, err)
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_USERID_BYLOTTERYID_ERROR), "FindOne, lotteryId:%v, error: %v", lotteryId, err)
	}
}

// NewLotteryModel returns a model for the database table.
func NewLotteryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LotteryModel {
	return &customLotteryModel{
		defaultLotteryModel: newLotteryModel(conn, c, opts...),
	}
}

func (c *customLotteryModel) GetLotterysByLessThanCurrentTime(ctx context.Context, currentTime time.Time, announceType int64) ([]int64, error) {
	var resp []int64
	query := fmt.Sprintf("SELECT id FROM %s WHERE announce_type = ? AND is_announced = 0 AND announce_time <= ?", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, announceType, currentTime)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GETLOTTERY_BYLESSTHAN_CURRENTTIME_ERROR), "GetLotterysByLessThanCurrentTime, CurrentTime:%v, anounceType:%v, error: %v", currentTime, announceType, err)
	}
	return resp, nil
}

// UpdateLotteryStatus 根据lotteryId更新lottery状态为已开奖
func (c *customLotteryModel) UpdateLotteryStatus(ctx context.Context, lotteryId int64) error {
	// 准备更新数据的SQL语句
	query := fmt.Sprintf("UPDATE %s SET is_announced = 1 WHERE id = ?", c.table)

	_, err := c.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.Exec(query, lotteryId)
	})
	if err != nil {
		return errors.Wrapf(xerr.NewErrCode(xerr.UPDATE_LOTTERY_STATUS_ERROR), "UpdateLotteryStatus, lotteryId:%v error: %v", lotteryId, err)
	}
	return nil
}

func (c *customLotteryModel) GetTypeIs2AndIsNotAnnounceLotterys(ctx context.Context, announceType int64) ([]*Lottery, error) {
	var resp []*Lottery
	query := fmt.Sprintf("SELECT * FROM %s WHERE announce_type = ? AND is_announced = 0", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, announceType)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.GET_TYPEIS2_AND_ISNOT_ANNOUNCE_LOTTERYS_ERROR), "GetTypeIs2AndIsNotAnnounceLotterys,announceType:%v, error: %v", announceType, err)
	}
	return resp, nil
}

func (c *customLotteryModel) GetLotteryIdByUserId(ctx context.Context, UserId int64) (*int64, error) {
	//func(ctx context.Context, conn sqlx.SqlConn, v any) error {
	query := fmt.Sprintf("select id from %s where user_id = ?", c.table)
	//	return conn.QueryRowCtx(ctx, v, query, UserId)
	//}
	var resp int64
	err := c.QueryRowNoCacheCtx(ctx, &resp, query, UserId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		//errors.Wrapf(xerr.NewErrCode(xerr.DB_LOTTERYID_NOTFOUND), "GetLotteryIdByUserId, UserId:%v, error: %v", UserId, err)
		return nil, nil
	default:
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_LOTTERYID_BYUSERID_ERROR), "FindOne, UserId:%v, error: %v", UserId, err)
	}
}

func (c *customLotteryModel) GetTodayLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error) {
	var resp []int64
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id = ? AND DATE(publish_time) = CURDATE()", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_WEEKLOTTERYIDS_BYUSREID_ERROR), "GetTodayLotteryIdsByUserId, user_id:%v, error: %v", UserId, err)
	}
	return resp, nil
}

func (c *customLotteryModel) GetWeekLotteryIdsByUserId(ctx context.Context, UserId int64) ([]int64, error) {
	var resp []int64
	query := fmt.Sprintf("SELECT id FROM %s WHERE user_id = ? AND YEARWEEK(publish_time) = YEARWEEK(CURDATE())", c.table)
	err := c.QueryRowsNoCacheCtx(ctx, &resp, query, UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_TODAYLOTTERYIDSBYUSERID_ERROR), "GetWeekLotteryIdsByUserId, user_id:%v, error: %v", UserId, err)
	}
	return resp, nil
}
