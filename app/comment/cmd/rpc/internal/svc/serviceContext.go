package svc

import (
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
		Config: c,
	}
}
