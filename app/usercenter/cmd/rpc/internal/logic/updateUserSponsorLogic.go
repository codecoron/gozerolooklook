package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserSponsorLogic {
	return &UpdateUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserSponsorLogic) UpdateUserSponsor(in *pb.UpdateUserSponsorReq) (*pb.UpdateUserSponsorResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserSponsorResp{}, nil
}
