package goodsInfo

import (
	"context"

	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLogic) Query(req *types.GoodsInfoReq) (resp *types.GoodsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
