// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	goodsInfo "looklook/app/shop/cmd/api/desc/internal/handler/goodsInfo"
	"looklook/app/shop/cmd/api/desc/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/goodsInfo/query",
				Handler: goodsInfo.QueryHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/shop/v1"),
	)
}
