package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserShopByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserShopByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserShopByIdLogic {
	return &GetUserShopByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserShopByIdLogic) GetUserShopById(in *pb.GetUserShopByIdReq) (*pb.GetUserShopByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserShopByIdResp{}, nil
}
