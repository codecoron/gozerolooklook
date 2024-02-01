package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchVoteConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchVoteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchVoteConfigLogic {
	return &SearchVoteConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchVoteConfigLogic) SearchVoteConfig(in *pb.SearchVoteConfigReq) (*pb.SearchVoteConfigResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchVoteConfigResp{}, nil
}
