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

	//根据页码查询商品信息列表
	list, err := l.svcCtx.ShopRpc.GetGoodsList(l.ctx, &shop.GoodsListReq{
		PageSize: req.PageSize, //传入页码
	})

	if err != nil {
		return nil, err
	}
	var goodsList []types.GoodsInfo   //创建用于存储商品信息的数组
	for _, item := range list.Goods { //遍历list中的Goods数组
		var temp types.GoodsInfo            //创建商品信息变量
		_ = copier.Copy(&temp, item)        //进行商品信息的拷贝
		goodsList = append(goodsList, temp) //再添加到存储商品信息的数组goodsList中
	}
	//最后返回存储商品信息的数组list即可
	return &types.GoodsListResp{List: goodsList}, nil
}
