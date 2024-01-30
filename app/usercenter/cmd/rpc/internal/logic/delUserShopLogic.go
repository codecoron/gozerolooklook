package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserShopLogic {
	return &DelUserShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserShopLogic) DelUserShop(in *pb.DelUserShopReq) (*pb.DelUserShopResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelUserShopResp{}, nil
}
