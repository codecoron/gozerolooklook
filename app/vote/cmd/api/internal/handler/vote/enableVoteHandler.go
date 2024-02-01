package vote

import (
	"looklook/app/vote/cmd/api/internal/handler/validator"
	"net/http"

	"looklook/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/app/vote/cmd/api/internal/logic/vote"
	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"
)

func EnableVoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EnableVoteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		validateErr := validator.Validate(&req)
		if validateErr != nil {
			result.ParamErrorResult(r, w, validateErr)
			return
		}

		l := vote.NewEnableVoteLogic(r.Context(), svcCtx)
		resp, err := l.EnableVote(&req)

		result.HttpResult(r, w, resp, err)
	}
}
