package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/shop/cmd/rpc/internal/svc"
	"looklook/app/shop/cmd/rpc/pb"
	"looklook/app/shop/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsByIdLogic {
	return &GetGoodsByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGoodsByIdLogic) GetGoodsById(in *pb.GoodsReq) (*pb.GoodsResp, error) {
	goods, err := l.svcCtx.GoodsModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "lottery_id:%d,err:%v", in.Id, err)
	}
	fmt.Println("商品信息为: ", goods)
	pbGoods := new(pb.Goods)
	_ = copier.Copy(pbGoods, goods)
	fmt.Println("pbGoods:", pbGoods)
	return (*pb.GoodsResp)(pbGoods), nil
}
