package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelVoteRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelVoteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelVoteRecordLogic {
	return &DelVoteRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelVoteRecordLogic) DelVoteRecord(in *pb.DelVoteRecordReq) (*pb.DelVoteRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelVoteRecordResp{}, nil
}
