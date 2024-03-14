package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/lottery"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLastIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLastIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLastIdLogic {
	return &GetLastIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLastIdLogic) GetLastId(req *types.GetLastIdReq) (resp *types.GetLastIdResp, err error) {
	id, err := l.svcCtx.LotteryRpc.GetLotteryListLastId(l.ctx, &lottery.GetLotteryListLastIdReq{})
	if err != nil {
		return nil, err
	}
	resp = &types.GetLastIdResp{
		LastId: id.LastId,
	}
	return
}
