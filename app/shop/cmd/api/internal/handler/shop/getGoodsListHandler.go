package shop

import (
	"looklook/app/shop/cmd/api/internal/handler/translator"
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/shop/cmd/api/internal/logic/shop"
	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"
)

func GetGoodsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := shop.NewGetGoodsListLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodsList(&req)

		result.HttpResult(r, w, resp, err)
	}
}
