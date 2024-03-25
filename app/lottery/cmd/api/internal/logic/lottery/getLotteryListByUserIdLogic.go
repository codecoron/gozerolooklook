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
		UserId:      userId,
		LastId:      req.LastId,
		Size:        req.Size,
		Type:        req.Type,
		IsAnnounced: req.IsAnnounced,
	})
	if err != nil {
		return nil, err
	}

	// 数据传递
	resp = &types.GetLotteryListByUserIdResp{}
	for _, v := range res.LotteryPrizes {
		var item types.LotteryPrizes
		item.LotteryId = v.LotteryId
		item.Time = v.Time
		item.Prizes = make([]*types.Prize, 0)
		for _, prize := range v.Prizes {
			item.Prizes = append(item.Prizes, &types.Prize{
				Id:        prize.Id,
				LotteryId: prize.LotteryId,
				CreatePrize: types.CreatePrize{
					Type:      prize.Type,
					Name:      prize.Name,
					Count:     prize.Count,
					Thumb:     prize.Thumb,
					Level:     prize.Level,
					GrantType: prize.GrantType,
				},
			})
		}
		resp.List = append(resp.List, item)
	}
	return
}
