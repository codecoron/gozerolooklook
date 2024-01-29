package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoteRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoteRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoteRecordByIdLogic {
	return &GetVoteRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVoteRecordByIdLogic) GetVoteRecordById(in *pb.GetVoteRecordByIdReq) (*pb.GetVoteRecordByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetVoteRecordByIdResp{}, nil
}
