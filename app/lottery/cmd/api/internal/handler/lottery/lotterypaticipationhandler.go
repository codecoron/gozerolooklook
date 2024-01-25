package lottery

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/lottery/cmd/api/internal/logic/lottery"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
)

func LotteryParticipationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LotteryParticipationReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := lottery.NewLotteryParticipationLogic(r.Context(), svcCtx)
		resp, err := l.LotteryParticipation(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
