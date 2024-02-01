package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

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
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get LotteryDetail"), "Failed to get SearchLottery err : %v ,req:%+v", err, req)
	}
	resp = new(types.LotteryDetailResp)
	_ = copier.Copy(resp, res)
	_ = copier.Copy(resp, res.Lottery)

	// 获取赞助商信息
	res2, err := l.svcCtx.LotteryRpc.LotterySponsor(l.ctx, &lottery.LotterySponsorReq{
		Id: req.Id,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get LotteryDetail"), "Failed to get SearchLottery err : %v ,req:%+v", err, req)
	}
	resp.Sponsor = new(types.LotterySponsor)
	_ = copier.Copy(resp.Sponsor, res2)
	return
}
