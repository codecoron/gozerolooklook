package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAddressLogic {
	return &UpdateUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserAddressLogic) UpdateUserAddress(in *pb.UpdateUserAddressReq) (*pb.UpdateUserAddressResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserAddressResp{}, nil
}
