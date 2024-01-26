package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchParticipationLogic {
	return &SearchParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchParticipationLogic) SearchParticipation(req *types.SearchLotteryParticipationReq) (resp *types.SearchLotteryParticipationResp, err error) {
	// todo: add your logic here and delete this line

	return
}
