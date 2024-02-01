package lottery

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"
)

type SetLotteryIsSelectedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetLotteryIsSelectedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLotteryIsSelectedLogic {
	return &SetLotteryIsSelectedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetLotteryIsSelectedLogic) SetLotteryIsSelected(req *types.SetLotteryIsSelectedReq) (resp *types.SetLotteryIsSelectedResp, err error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	selectedLottery, err := l.svcCtx.LotteryRpc.SetIsSelectedLottery(l.ctx, &lottery.SetIsSelectedLotteryReq{
		Id:     req.Id,
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	return &types.SetLotteryIsSelectedResp{IsSelected: selectedLottery.IsSelected}, nil
}
