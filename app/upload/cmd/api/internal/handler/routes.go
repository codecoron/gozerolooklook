// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	upload "looklook/app/upload/cmd/api/internal/handler/upload"
	"looklook/app/upload/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/upload/add",
				Handler: upload.UploadHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/upload/v1"),
	)
}
