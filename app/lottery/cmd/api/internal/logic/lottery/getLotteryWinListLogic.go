package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryWinListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLotteryWinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryWinListLogic {
	return &GetLotteryWinListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLotteryWinListLogic) GetLotteryWinList(req *types.GetLotteryWinListReq) (resp *types.GetLotteryWinListResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	list, err := l.svcCtx.LotteryRpc.GetWonList(l.ctx, &lottery.GetWonListReq{
		UserId: userId,
		Page:   req.Page,
		Size:   req.Size,
		LastId: req.LastId,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.GetLotteryWinListResp)
	resp.List = make([]*types.WonList, len(list.List))
	for i, item := range list.List {
		resp.List[i] = &types.WonList{
			Id:        item.Id,
			LotteryId: item.LotteryId,
			UserId:    item.UserId,
			IsWon:     1,
			Prize: &types.Prizes{
				Id:        item.Prize.Id,
				LotteryId: item.Prize.LotteryId,
				Type:      item.Prize.Type,
				Name:      item.Prize.Name,
				Thumb:     item.Prize.Thumb,
				Count:     item.Prize.Count,
				GrantType: item.Prize.GrantType,
			},
		}
	}

	return
}
