package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTasksLogic {
	return &DelTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTasksLogic) DelTasks(in *pb.DelTasksReq) (*pb.DelTasksResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelTasksResp{}, nil
}
