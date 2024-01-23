package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UploadFileInfoModel = (*customUploadFileInfoModel)(nil)

type (
	// UploadFileInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUploadFileInfoModel.
	UploadFileInfoModel interface {
		uploadFileInfoModel
	}

	customUploadFileInfoModel struct {
		*defaultUploadFileInfoModel
	}
)

// NewUploadFileInfoModel returns a model for the database table.
func NewUploadFileInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UploadFileInfoModel {
	return &customUploadFileInfoModel{
		defaultUploadFileInfoModel: newUploadFileInfoModel(conn, c, opts...),
	}
}
