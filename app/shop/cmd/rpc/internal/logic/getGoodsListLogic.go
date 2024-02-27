package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/shop/cmd/rpc/internal/svc"
	"looklook/app/shop/cmd/rpc/pb"
	"looklook/app/shop/model"
	"looklook/common/xerr"
)

type GetGoodsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsListLogic {
	return &GetGoodsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsListLogic) GetGoodsList(in *pb.GoodsListReq) (*pb.GoodsListResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.GoodsModel.List(l.ctx, in.PageSize, 5)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "PageSize:%d,err:%v", in.PageSize, err)
	}
	goods := make([]*pb.Goods, 0)
	for _, goodsInfo := range list {
		pbGoods := new(pb.Goods)
		err := copier.Copy(pbGoods, goodsInfo)
		if err != nil {
			return nil, err
		}
		goods = append(goods, pbGoods)
	}
	return &pb.GoodsListResp{
		Goods: goods,
	}, nil
}
