package lottery

import (
	"context"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCreateLotteryListByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCreateLotteryListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCreateLotteryListByUserIdLogic {
	return &GetCreateLotteryListByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCreateLotteryListByUserIdLogic) GetCreateLotteryListByUserId(req *types.GetCreateLotteryListByUserIdReq) (resp *types.GetCreateLotteryListByUserIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
