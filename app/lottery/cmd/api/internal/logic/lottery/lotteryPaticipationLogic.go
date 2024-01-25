package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryParticipationLogic {
	return &LotteryParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryParticipationLogic) LotteryParticipation(req *types.LotteryParticipationReq) (resp *types.LotteryParticipationResp, err error) {
	// todo: add your logic here and delete this line

	return
}
