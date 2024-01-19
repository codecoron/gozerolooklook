package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/lottery"
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

func (l *LotteryDetailLogic) LotteryDetail(req *types.LotteryDetailReq) (*types.LotteryDetailResp, error) {
	resp, err := l.svcCtx.LotteryRpc.LotteryDetail(l.ctx, &lottery.LotteryDetailReq{
		Id: req.Id,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get LotteryDetail"), "Failed to get SearchLottery err : %v ,req:%+v", err, req)
	}
	var lotteryDetail types.LotteryDetail
	_ = copier.Copy(&lotteryDetail, resp)
	lotteryDetail.Id = resp.Lottery.Id
	lotteryDetail.Name = resp.Lottery.Name
	lotteryDetail.IsSelected = resp.Lottery.IsSelected
	lotteryDetail.UserId = resp.Lottery.UserId
	lotteryDetail.AwardDeadline = resp.Lottery.AwardDeadline
	lotteryDetail.Introduce = resp.Lottery.Introduce
	lotteryDetail.JoinNumber = resp.Lottery.JoinNumber
	lotteryDetail.PublishTime = resp.Lottery.PublishTime
	lotteryDetail.PublishType = resp.Lottery.PublishType

	return &types.LotteryDetailResp{LotteryDetail: lotteryDetail}, nil
}
