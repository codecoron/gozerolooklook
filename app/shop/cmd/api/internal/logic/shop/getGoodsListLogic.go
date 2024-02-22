package shop

import (
	"context"

	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsListLogic) GetGoodsList(req *types.GoodsListReq) (resp *types.GoodsListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
