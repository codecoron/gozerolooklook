package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserFollowLogic {
	return &DelUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserFollowLogic) DelUserFollow(in *pb.DelUserFollowReq) (*pb.DelUserFollowResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelUserFollowResp{}, nil
}
