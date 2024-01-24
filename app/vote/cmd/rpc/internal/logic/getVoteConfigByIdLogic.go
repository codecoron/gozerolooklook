package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoteConfigByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoteConfigByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoteConfigByIdLogic {
	return &GetVoteConfigByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVoteConfigByIdLogic) GetVoteConfigById(in *pb.GetVoteConfigByIdReq) (*pb.GetVoteConfigByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetVoteConfigByIdResp{}, nil
}
