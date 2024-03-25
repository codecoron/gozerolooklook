package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserDynamicByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserDynamicByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserDynamicByIdLogic {
	return &GetUserDynamicByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserDynamicByIdLogic) GetUserDynamicById(in *pb.GetUserDynamicByIdReq) (*pb.GetUserDynamicByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserDynamicByIdResp{}, nil
}
