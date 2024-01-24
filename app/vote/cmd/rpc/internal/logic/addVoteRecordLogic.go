package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVoteRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVoteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVoteRecordLogic {
	return &AddVoteRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------投票记录表-----------------------
func (l *AddVoteRecordLogic) AddVoteRecord(in *pb.AddVoteRecordReq) (*pb.AddVoteRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddVoteRecordResp{}, nil
}
