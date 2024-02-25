// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "looklook/app/comment/cmd/api/internal/handler/comment"
	"looklook/app/comment/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comment/Test",
				Handler: comment.TestHandler(serverCtx),
			},
		},
		rest.WithPrefix("/comment/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comment/addComment",
				Handler: comment.AddCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/deleteComment",
				Handler: comment.DeleteCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/updateComment",
				Handler: comment.UpdateCommentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/commentPraise",
				Handler: comment.CommentPraiseHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/comment/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/comment/getCommentList",
				Handler: comment.GetCommentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/getCommentDetail",
				Handler: comment.GetCommentDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/comment/getLastId",
				Handler: comment.GetCommentLastIdHandler(serverCtx),
			},
		},
		rest.WithPrefix("/comment/v1"),
	)
}
