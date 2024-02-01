package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTasksLogic {
	return &UpdateTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTasksLogic) UpdateTasks(in *pb.UpdateTasksReq) (*pb.UpdateTasksResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateTasksResp{}, nil
}
