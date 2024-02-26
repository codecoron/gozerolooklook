package shop

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"
	"looklook/app/shop/cmd/rpc/shop"

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
	res, err := l.svcCtx.ShopRpc.GetGoodsById(l.ctx, &shop.GoodsReq{
		Id: req.GoodsId,
	})
	if err != nil {
		return nil, err
	}

	resp = new(types.GoodsInfoResp)
	_ = copier.Copy(resp.GoodsInfo, res.Goods)

	return resp, nil
}
