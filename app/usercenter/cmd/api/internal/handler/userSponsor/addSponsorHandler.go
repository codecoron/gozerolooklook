package userSponsor

import (
	"looklook/app/usercenter/cmd/api/internal/handler/translator"
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/usercenter/cmd/api/internal/logic/userSponsor"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
)

func AddSponsorHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateSponosorReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := userSponsor.NewAddSponsorLogic(r.Context(), svcCtx)
		resp, err := l.AddSponsor(&req)

		result.HttpResult(r, w, resp, err)
	}
}
