package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ NoticeSubscribeSettingModel = (*customNoticeSubscribeSettingModel)(nil)

type (
	// NoticeSubscribeSettingModel is an interface to be customized, add more methods here,
	// and implement the added methods in customNoticeSubscribeSettingModel.
	NoticeSubscribeSettingModel interface {
		noticeSubscribeSettingModel
	}

	customNoticeSubscribeSettingModel struct {
		*defaultNoticeSubscribeSettingModel
	}
)

// NewNoticeSubscribeSettingModel returns a model for the database table.
func NewNoticeSubscribeSettingModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) NoticeSubscribeSettingModel {
	return &customNoticeSubscribeSettingModel{
		defaultNoticeSubscribeSettingModel: newNoticeSubscribeSettingModel(conn, c, opts...),
	}
}
