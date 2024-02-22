package shop

import (
	"context"

	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsByIdLogic {
	return &GetGoodsByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsByIdLogic) GetGoodsById(req *types.GoodsInfoReq) (resp *types.GoodsInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
