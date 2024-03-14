package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLotteryLogic {
	return &PublishLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLotteryLogic) PublishLottery(req *types.PublishLotteryReq) (resp *types.PublishLotteryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
