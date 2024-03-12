package userWonDynamicComment

import (
	"looklook/app/usercenter/cmd/api/internal/handler/translator"
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/usercenter/cmd/api/internal/logic/userWonDynamicComment"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
)

func UserWonDynamicCommentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserWonDynamicCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := userWonDynamicComment.NewUserWonDynamicCommentListLogic(r.Context(), svcCtx)
		resp, err := l.UserWonDynamicCommentList(&req)

		result.HttpResult(r, w, resp, err)
	}
}
