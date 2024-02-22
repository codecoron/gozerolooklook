package checkin

import (
	"looklook/app/checkin/cmd/api/internal/handler/translator"
	"looklook/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/checkin/cmd/api/internal/logic/checkin"
	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"
)

func GetCheckinHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCheckinReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := translator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := checkin.NewGetCheckinLogic(r.Context(), svcCtx)
		resp, err := l.GetCheckin(&req)
		//注意 handler这里需要用result.HttpResult() 才会返回    "code": 200, "msg": "OK",
		result.HttpResult(r, w, resp, err)
	}
}
