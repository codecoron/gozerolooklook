package userSponsor

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/usercenter/cmd/api/internal/logic/userSponsor"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
)

func SponsorDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SponosorDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := userSponsor.NewSponsorDetailLogic(r.Context(), svcCtx)
		resp, err := l.SponsorDetail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
