package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserContactLogic {
	return &UpdateUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserContactLogic) UpdateUserContact(in *pb.UpdateUserContactReq) (*pb.UpdateUserContactResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateUserContactResp{}, nil
}
