package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDynamicByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDynamicByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDynamicByUserIdLogic {
	return &GetUserDynamicByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserDynamicByUserIdLogic) GetUserDynamicByUserId(in *pb.GetUserDynamicByUserIdReq) (*pb.GetUserDynamicByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserDynamicByUserIdResp{}, nil
}
