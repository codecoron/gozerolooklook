package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/comment/cmd/rpc/internal/config"
	"looklook/app/comment/model"
)

type ServiceContext struct {
	Config       config.Config
	CommentModel model.CommentModel
	PraiseModel  model.PraiseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		CommentModel: model.NewCommentModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		PraiseModel:  model.NewPraiseModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
