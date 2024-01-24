package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVoteConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVoteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVoteConfigLogic {
	return &AddVoteConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------投票表-----------------------
func (l *AddVoteConfigLogic) AddVoteConfig(in *pb.AddVoteConfigReq) (*pb.AddVoteConfigResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddVoteConfigResp{}, nil
}
