package logic

import (
	"context"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVoteRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVoteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVoteRecordLogic {
	return &UpdateVoteRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVoteRecordLogic) UpdateVoteRecord(in *pb.UpdateVoteRecordReq) (*pb.UpdateVoteRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateVoteRecordResp{}, nil
}
