package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSponsorByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSponsorByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSponsorByIdLogic {
	return &GetUserSponsorByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserSponsorByIdLogic) GetUserSponsorById(in *pb.GetUserSponsorByIdReq) (*pb.GetUserSponsorByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetUserSponsorByIdResp{}, nil
}
