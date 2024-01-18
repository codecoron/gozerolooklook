package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAddressByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressByIdLogic {
	return &GetUserAddressByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserAddressByIdLogic) GetUserAddressById(in *pb.GetUserAddressByIdReq) (*pb.GetUserAddressByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserAddressByIdResp{}, nil
}
