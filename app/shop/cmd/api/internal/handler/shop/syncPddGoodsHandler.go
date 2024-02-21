package shop

import (
	"net/http"

	"looklook/common/result"

	"looklook/app/shop/cmd/api/internal/logic/shop"
	"looklook/app/shop/cmd/api/internal/svc"
)

func SyncPddGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := shop.NewSyncPddGoodsLogic(r.Context(), svcCtx)
		err := l.SyncPddGoods()

		result.HttpResult(r, w, nil, err)
	}
}
