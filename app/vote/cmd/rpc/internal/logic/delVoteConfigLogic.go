package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelVoteConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelVoteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelVoteConfigLogic {
	return &DelVoteConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelVoteConfigLogic) DelVoteConfig(in *pb.DelVoteConfigReq) (*pb.DelVoteConfigResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelVoteConfigResp{}, nil
}
