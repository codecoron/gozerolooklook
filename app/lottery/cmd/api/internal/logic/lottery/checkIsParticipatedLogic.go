package lottery

import (
	"context"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsParticipatedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckIsParticipatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsParticipatedLogic {
	return &CheckIsParticipatedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckIsParticipatedLogic) CheckIsParticipated(req *types.CheckIsParticipatedReq) (resp *types.CheckIsParticipatedResp, err error) {
	// 需要获取当前用户id，从而判断当前用户是否有参与当前lottery
	userId := ctxdata.GetUidFromCtx(l.ctx)
	participated, err := l.svcCtx.LotteryRpc.CheckIsParticipated(l.ctx, &lottery.CheckIsParticipatedReq{
		UserId:    userId,
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CheckIsParticipatedResp{IsParticipated: participated.IsParticipated}, nil
}
