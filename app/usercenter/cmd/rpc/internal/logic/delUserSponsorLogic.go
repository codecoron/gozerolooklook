package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserSponsorLogic {
	return &DelUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserSponsorLogic) DelUserSponsor(in *pb.DelUserSponsorReq) (*pb.DelUserSponsorResp, error) {
	// todo: add your logic here and delete this line
	l.svcCtx.UserSponsorModel.Delete(l.ctx, in.Id)

	return &pb.DelUserSponsorResp{}, nil
}
