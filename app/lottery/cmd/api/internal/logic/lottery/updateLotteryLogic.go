package lottery

import (
	"context"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLotteryLogic {
	return &UpdateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLotteryLogic) UpdateLottery(req *types.UpdateLotteryReq) (resp *types.UpdateLotteryResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.LotteryRpc.UpdateLottery(l.ctx, &lottery.UpdateLotteryReq{
		UserId: userId,
		Id:     req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateLotteryResp{}, nil
}
