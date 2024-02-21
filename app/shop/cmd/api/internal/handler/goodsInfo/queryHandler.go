package goodsInfo

import (
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/shop/cmd/api/internal/logic/goodsInfo"
	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"
)

func QueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		//validateErr := translator.Validate(&req)
		//if validateErr != nil {
		//	result.ParamErrorResult(r, w, validateErr)
		//	return
		//}

		l := goodsInfo.NewQueryLogic(r.Context(), svcCtx)
		resp, err := l.Query(&req)

		result.HttpResult(r, w, resp, err)
	}
}
