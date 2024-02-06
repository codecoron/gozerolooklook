package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserContactByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserContactByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserContactByIdLogic {
	return &GetUserContactByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserContactByIdLogic) GetUserContactById(in *pb.GetUserContactByIdReq) (*pb.GetUserContactByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserContactByIdResp{}, nil
}
