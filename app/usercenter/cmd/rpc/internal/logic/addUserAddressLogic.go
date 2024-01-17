package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserAddressLogic {
	return &AddUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户收货地址表-----------------------
func (l *AddUserAddressLogic) AddUserAddress(in *pb.AddUserAddressReq) (*pb.AddUserAddressResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddUserAddressResp{}, nil
}
