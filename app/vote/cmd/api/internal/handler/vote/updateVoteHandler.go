package vote

import (
	"looklook/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/vote/cmd/api/internal/logic/vote"
	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"
)

func UpdateVoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateVoteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := vote.NewUpdateVoteLogic(r.Context(), svcCtx)
		resp, err := l.UpdateVote(&req)

		result.HttpResult(r, w, resp, err)
	}
}
