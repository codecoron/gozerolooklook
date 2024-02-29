package lottery

import (
	"looklook/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/lottery/cmd/api/internal/logic/lottery"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
)

func ChanceTypeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChanceTypeListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := lottery.NewChanceTypeListLogic(r.Context(), svcCtx)
		resp, err := l.ChanceTypeList(&req)
		result.HttpResult(r, w, resp, err)
	}
}
