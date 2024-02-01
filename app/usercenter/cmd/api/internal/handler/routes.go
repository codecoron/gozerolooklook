// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	address "looklook/app/usercenter/cmd/api/internal/handler/address"
	user "looklook/app/usercenter/cmd/api/internal/handler/user"
	userContact "looklook/app/usercenter/cmd/api/internal/handler/userContact"
	userfollow "looklook/app/usercenter/cmd/api/internal/handler/userfollow"
	"looklook/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/addAddress",
				Handler: address.AddAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/addressList",
				Handler: address.AddressListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/convertAddress",
				Handler: address.ConvertAddressHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/wxMiniAuth",
				Handler: user.WxMiniAuthHandler(serverCtx),
			},
		},
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/detail",
				Handler: user.DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/setAdmin",
				Handler: user.SetAdminHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/user/update",
				Handler: user.UpdateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/userContact/create",
				Handler: userContact.CreateHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/userfollow/addfollowInfo",
				Handler: userfollow.AddfollowInfoHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/usercenter/v1"),
	)
}
