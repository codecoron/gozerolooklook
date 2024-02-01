package logic

import (
	"context"
	"github.com/pkg/errors"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetAdminLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAdminLogic {
	return &SetAdminLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetAdminLogic) SetAdmin(in *pb.SetAdminReq) (*pb.SetAdminResp, error) {
	err := l.svcCtx.UserModel.SetAdmin(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", in)
	}
	return &pb.SetAdminResp{}, nil
}
