package shop

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/shop/cmd/api/internal/svc"
	"looklook/app/shop/cmd/api/internal/types"
	"looklook/app/shop/cmd/rpc/shop"
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
		Id: req.Id,
	})
	fmt.Println(res)
	logx.Info(res)
	if err != nil {
		return nil, err
	}

	resp = new(types.GoodsInfoResp)
	_ = copier.Copy(resp, res)
	return resp, nil
}
