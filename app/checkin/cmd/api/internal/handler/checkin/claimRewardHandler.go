package checkin

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/checkin/cmd/api/internal/logic/checkin"
	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"
)

func ClaimRewardHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClaimRewardReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := checkin.NewClaimRewardLogic(r.Context(), svcCtx)
		resp, err := l.ClaimReward(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
