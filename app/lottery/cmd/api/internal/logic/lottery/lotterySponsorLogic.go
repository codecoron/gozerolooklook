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

type LotterySponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLotterySponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotterySponsorLogic {
	return &LotterySponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LotterySponsorLogic) LotterySponsor(req *types.LotterySponsorReq) (*types.LotterySponsorResp, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.LotteryRpc.LotterySponsor(l.ctx, &lottery.LotterySponsorReq{
		Id: req.Id,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get LotteryDetail"), "Failed to get SearchLottery err : %v ,req:%+v", err, req)
	}
	var sponsor types.Sponsor
	_ = copier.Copy(&sponsor, resp)
	return &types.LotterySponsorResp{Sponsor: sponsor}, nil
}
