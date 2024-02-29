package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserContactLogic {
	return &EditUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditUserContactLogic) EditUserContact(in *pb.EditUserContactReq) (*pb.EditUserContactResp, error) {
	// todo: add your logic here and delete this line

	return &pb.EditUserContactResp{}, nil
}
