package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLotteryParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryParticipationLogic {
	return &AddLotteryParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLotteryParticipationLogic) AddLotteryParticipation(req *types.AddLotteryParticipationReq) (resp *types.AddLotteryParticipationResp, err error) {
	_, err = l.svcCtx.LotteryRpc.AddLotteryParticipation(l.ctx, &lottery.AddLotteryParticipationReq{
		UserId:    ctxdata.GetUidFromCtx(l.ctx),
		LotteryId: req.LotteryId,
	})
	return
}
