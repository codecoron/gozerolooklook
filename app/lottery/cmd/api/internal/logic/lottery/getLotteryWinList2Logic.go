package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/lottery"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryWinList2Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLotteryWinList2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryWinList2Logic {
	return &GetLotteryWinList2Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLotteryWinList2Logic) GetLotteryWinList2(req *types.GetLotteryWinList2Req) (resp *types.GetLotteryWinList2Resp, err error) {
	list, err := l.svcCtx.LotteryRpc.GetWonListByLotteryId(l.ctx, &lottery.GetWonListByLotteryIdReq{
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}

	// 数据传递
	resp = &types.GetLotteryWinList2Resp{}
	for _, v := range list.List {
		var item types.WonList2
		err = copier.Copy(&item, &v)
		if err != nil {
			return nil, err
		}
		resp.List = append(resp.List, &item)
	}
	if err != nil {
		return nil, err
	}

	return
}
