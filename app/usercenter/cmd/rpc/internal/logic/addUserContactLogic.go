package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserContactLogic {
	return &AddUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------抽奖发起人联系方式-----------------------
func (l *AddUserContactLogic) AddUserContact(in *pb.AddUserContactReq) (*pb.AddUserContactResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddUserContactResp{}, nil
}
