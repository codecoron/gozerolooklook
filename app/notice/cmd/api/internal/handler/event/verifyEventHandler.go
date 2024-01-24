package event

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/notice/cmd/api/internal/logic/event"
	"looklook/app/notice/cmd/api/internal/svc"
	"looklook/app/notice/cmd/api/internal/types"
)

func VerifyEventHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerifyEventReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := event.NewVerifyEventLogic(r.Context(), svcCtx)
		resp, err := l.VerifyEvent(&req, w)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
