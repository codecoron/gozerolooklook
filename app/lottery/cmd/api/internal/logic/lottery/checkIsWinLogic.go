package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsWinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckIsWinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsWinLogic {
	return &CheckIsWinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckIsWinLogic) CheckIsWin(req *types.CheckIsWinReq) (resp *types.CheckIsWinResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	isWon, err := l.svcCtx.LotteryRpc.CheckUserIsWon(l.ctx, &lottery.CheckUserIsWonReq{
		UserId:    userId,
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CheckIsWinResp{IsWon: isWon.IsWon}, nil
}
