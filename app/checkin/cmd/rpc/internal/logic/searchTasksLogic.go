package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTasksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTasksLogic {
	return &SearchTasksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTasksLogic) SearchTasks(in *pb.SearchTasksReq) (*pb.SearchTasksResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchTasksResp{}, nil
}
