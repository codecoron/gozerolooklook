package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"
)

type GetTaskProgressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskProgressLogic {
	return &GetTaskProgressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskProgressLogic) GetTaskProgress(in *pb.GetTaskProgressReq) (*pb.GetTaskProgressResp, error) {
	// todo: 调用其他服务查询任务进度

	return &pb.GetTaskProgressResp{}, nil
}
