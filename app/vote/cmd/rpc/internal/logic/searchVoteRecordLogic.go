package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchVoteRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchVoteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchVoteRecordLogic {
	return &SearchVoteRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchVoteRecordLogic) SearchVoteRecord(in *pb.SearchVoteRecordReq) (*pb.SearchVoteRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchVoteRecordResp{}, nil
}
