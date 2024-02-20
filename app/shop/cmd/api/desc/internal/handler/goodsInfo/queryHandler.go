package goodsInfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/shop/cmd/api/desc/internal/logic/goodsInfo"
	"looklook/app/shop/cmd/api/desc/internal/svc"
	"looklook/app/shop/cmd/api/desc/internal/types"
)

func QueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goodsInfo.NewQueryLogic(r.Context(), svcCtx)
		resp, err := l.Query(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
