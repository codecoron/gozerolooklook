package user

import (
	"looklook/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/usercenter/cmd/api/internal/logic/user"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
)

func SetAdminHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetAdminReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSetAdminLogic(r.Context(), svcCtx)
		resp, err := l.SetAdmin(&req)
		if err != nil {
			result.ParamErrorResult(r, w, err)
		} else {
			result.HttpResult(r, w, resp, err)
		}
	}
}
