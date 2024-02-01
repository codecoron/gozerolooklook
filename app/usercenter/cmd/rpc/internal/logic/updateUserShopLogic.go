package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserShopLogic {
	return &UpdateUserShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserShopLogic) UpdateUserShop(in *pb.UpdateUserShopReq) (*pb.UpdateUserShopResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserShopResp{}, nil
}
