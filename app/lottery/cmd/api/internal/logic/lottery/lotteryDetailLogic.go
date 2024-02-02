package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotteryDetailLogic) LotteryDetail(req *types.LotteryDetailReq) (resp *types.LotteryDetailResp, err error) {
	// 需要获取当前用户id，从而判断当前用户是否有参与当前lottery
	userId := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.LotteryRpc.LotteryDetail(l.ctx, &lottery.LotteryDetailReq{
		Id:     req.Id,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}
	resp = new(types.LotteryDetailResp)
	// todo 返回成功，但是json反序列化提示error
	_ = copier.Copy(resp, res)
	_ = copier.Copy(resp, res.Lottery)

	// todo 获取赞助商信息
	//res2, err := l.svcCtx.LotteryRpc.LotterySponsor(l.ctx, &lottery.LotterySponsorReq{
	//	Id: req.Id,
	//})
	//if err != nil {
	//	return nil, err
	//}
	//resp.Sponsor = new(types.LotterySponsor)
	//_ = copier.Copy(resp.Sponsor, res2)
	return
}
