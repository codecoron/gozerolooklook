package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTaskRecordLogic {
	return &UpdateTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTaskRecordLogic) UpdateTaskRecord(in *pb.UpdateTaskRecordReq) (*pb.UpdateTaskRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateTaskRecordResp{}, nil
}
