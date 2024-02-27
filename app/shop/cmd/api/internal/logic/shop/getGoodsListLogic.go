package shop

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/shop/cmd/rpc/shop"

	"github.com/zeromicro/go-zero/core/logx"

	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"
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

	list, err := l.svcCtx.ShopRpc.GetGoodsList(l.ctx, &shop.GoodsListReq{
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var goodsList []types.GoodsInfo
	for _, item := range list.Goods {
		var temp types.GoodsInfo
		_ = copier.Copy(&temp, item)
		goodsList = append(goodsList, temp)
	}
	return &types.GoodsListResp{List: goodsList}, nil
}
