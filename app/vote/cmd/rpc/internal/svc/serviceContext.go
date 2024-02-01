package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/vote/cmd/rpc/internal/config"
	"looklook/app/vote/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	VoteConfigModel    model.VoteConfigModel
	VoteConfigDiyModel model.VoteConfigDiyModel
	VoteRecordModel    model.VoteRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),

		VoteConfigModel: model.NewVoteConfigModel(sqlConn, c.Cache),
		VoteRecordModel: model.NewVoteRecordModel(sqlConn, c.Cache),
	}
}
