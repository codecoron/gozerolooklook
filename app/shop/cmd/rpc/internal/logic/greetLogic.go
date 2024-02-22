package logic

import (
	"context"

	"looklook/app/shop/cmd/rpc/internal/svc"
	"looklook/app/shop/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGreetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetLogic {
	return &GreetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GreetLogic) Greet(in *pb.GoodsReq) (*pb.GoodsResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GoodsResp{}, nil
}
