package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/app/usercenter/cmd/rpc/internal/config"
	"looklook/app/usercenter/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis

	UserModel        model.UserModel
	UserAuthModel    model.UserAuthModel
	UserAddressModel model.UserAddressModel
	UserSponsorModel model.UserSponsorModel
	UserContactModel model.UserContactModel
	CheckinRpc       checkin.Checkin
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),

		UserAuthModel:    model.NewUserAuthModel(sqlConn, c.Cache),
		UserModel:        model.NewUserModel(sqlConn, c.Cache),
		UserAddressModel: model.NewUserAddressModel(sqlConn, c.Cache),
		UserSponsorModel: model.NewUserSponsorModel(sqlConn, c.Cache),
		UserContactModel: model.NewUserContactModel(sqlConn, c.Cache),
		CheckinRpc:       checkin.NewCheckin(zrpc.MustNewClient(c.CheckinRpcConf)),
	}
}
