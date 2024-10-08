package comment

import (
	"looklook/app/comment/cmd/api/internal/handler/translator"
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/comment/cmd/api/internal/logic/comment"
	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"
)

func AddCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CommentAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := comment.NewAddCommentLogic(r.Context(), svcCtx)
		resp, err := l.AddComment(&req)

		result.HttpResult(r, w, resp, err)
	}
}
