package model

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ClockTaskRecordModel = (*customClockTaskRecordModel)(nil)

type (
	// ClockTaskRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customClockTaskRecordModel.
	ClockTaskRecordModel interface {
		clockTaskRecordModel
		// 自定义方法
		// 传入lotteryId以及userIds，得到每个参与者的中奖倍率
		GetClockTaskRecordByLotteryIdAndUserIds(lotteryId int64, userIds []int64) ([]*ClockTaskRecord, error)
	}

	customClockTaskRecordModel struct {
		*defaultClockTaskRecordModel
	}
)

// NewClockTaskRecordModel returns a model for the database table.
func NewClockTaskRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ClockTaskRecordModel {
	return &customClockTaskRecordModel{
		defaultClockTaskRecordModel: newClockTaskRecordModel(conn, c, opts...),
	}
}

func (m *defaultClockTaskRecordModel) GetClockTaskRecordByLotteryIdAndUserIds(lotteryId int64, userIds []int64) ([]*ClockTaskRecord, error) {
	if len(userIds) == 0 {
		return nil, nil
	}

	// 将userIds转换为字符串
	userIdsStr := ""
	for i, userId := range userIds {
		if i == 0 {
			userIdsStr = fmt.Sprintf("%d", userId)
		} else {
			userIdsStr = fmt.Sprintf("%s,%d", userIdsStr, userId)
		}
	}

	query := fmt.Sprintf("select %s from %s where lottery_id = ? and user_id in (%s)", clockTaskRecordRows, m.table, userIdsStr)
	var records []*ClockTaskRecord
	err := m.QueryRowsNoCache(&records, query, lotteryId)
	if err != nil {
		return nil, err
	}
	return records, nil
}
