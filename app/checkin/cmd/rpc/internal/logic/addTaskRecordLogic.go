package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTaskRecordLogic {
	return &AddTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------taskRecord-----------------------
func (l *AddTaskRecordLogic) AddTaskRecord(in *pb.AddTaskRecordReq) (*pb.AddTaskRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddTaskRecordResp{}, nil
}
