package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserAddressLogic {
	return &DelUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserAddressLogic) DelUserAddress(in *pb.DelUserAddressReq) (*pb.DelUserAddressResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelUserAddressResp{}, nil
}
