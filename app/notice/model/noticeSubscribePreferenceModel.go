package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NoticeSubscribePreferenceModel = (*customNoticeSubscribePreferenceModel)(nil)

type (
	// NoticeSubscribePreferenceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoticeSubscribePreferenceModel.
	NoticeSubscribePreferenceModel interface {
		noticeSubscribePreferenceModel
		Upsert(ctx context.Context, data *NoticeSubscribePreference) (sql.Result, error)
	}

	customNoticeSubscribePreferenceModel struct {
		*defaultNoticeSubscribePreferenceModel
	}
)

// NewNoticeSubscribePreferenceModel returns a model for the database table.
func NewNoticeSubscribePreferenceModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NoticeSubscribePreferenceModel {
	return &customNoticeSubscribePreferenceModel{
		defaultNoticeSubscribePreferenceModel: newNoticeSubscribePreferenceModel(conn, c, opts...),
	}
}

func (m *customNoticeSubscribePreferenceModel) Upsert(ctx context.Context, data *NoticeSubscribePreference) (sql.Result, error) {
	noticeNoticeSubscribePreferenceIdKey := fmt.Sprintf("%s%v", cacheNoticeNoticeSubscribePreferenceIdPrefix, data.Id)
	noticeNoticeSubscribePreferenceUserOpenidMsgTemplateIdKey := fmt.Sprintf("%s%v:%v", cacheNoticeNoticeSubscribePreferenceUserOpenidMsgTemplateIdPrefix, data.UserOpenid, data.MsgTemplateId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?) ON DUPLICATE KEY UPDATE accept_count = ?", m.table, noticeSubscribePreferenceRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserOpenid, data.MsgTemplateId, data.AcceptCount, data.AcceptCount)
	}, noticeNoticeSubscribePreferenceIdKey, noticeNoticeSubscribePreferenceUserOpenidMsgTemplateIdKey)
	return ret, err
}
