package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryListByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLotteryListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListByUserIdLogic {
	return &GetLotteryListByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLotteryListByUserIdLogic) GetLotteryListByUserId(req *types.GetLotteryListByUserIdReq) (resp *types.GetLotteryListByUserIdResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.LotteryRpc.GetLotteryPrizesListByUserId(l.ctx, &lottery.GetLotteryPrizesListByUserIdReq{
		UserId: userId,
		Page:   req.Page,
		Size:   req.Size,
		Type:   req.Type,
	})
	if err != nil {
		return nil, err
	}

	// 数据传递
	resp = &types.GetLotteryListByUserIdResp{}
	for _, v := range res.LotteryPrizes {
		var item types.LotteryPrizes
		item.LotteryId = v.LotteryId
		if req.Type == 1 {
			if v.CreateTime != 0 {
				item.Time = v.CreateTime
			} else {
				item.Time = v.ParticipateTime
			}
		} else if req.Type == 2 {
			item.Time = v.ParticipateTime
		} else if req.Type == 3 {
			item.Time = v.WonTime
		}
		item.Prizes = make([]*types.CreatePrize, 0)
		for _, prize := range v.Prizes {
			item.Prizes = append(item.Prizes, &types.CreatePrize{
				Type:      prize.Type,
				Name:      prize.Name,
				Count:     prize.Count,
				Level:     prize.Level,
				Thumb:     prize.Thumb,
				GrantType: prize.GrantType,
			})
		}
		resp.List = append(resp.List, item)
	}
	return
}
