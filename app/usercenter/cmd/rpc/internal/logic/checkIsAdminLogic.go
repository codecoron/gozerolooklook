package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsAdminLogic {
	return &CheckIsAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsAdminLogic) CheckIsAdmin(in *pb.CheckIsAdminReq) (*pb.CheckIsAdminResp, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	var isAdmin bool
	if user.IsAdmin == 1 {
		isAdmin = true
	} else {
		isAdmin = false
	}
	return &pb.CheckIsAdminResp{
		IsAdmin: isAdmin,
	}, nil
}
