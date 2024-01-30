package checkin

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/common/ctxdata"

	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTasksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTasksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTasksLogic {
	return &GetTasksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTasksLogic) GetTasks(req *types.GetTasksReq) (resp *types.GetTasksResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	tasks, err := l.svcCtx.CheckinRpc.GetTaskRecordByUserId(l.ctx, &checkin.GetTaskRecordByUserIdReq{
		UserId: userId,
	})
	//logx.Error("api,tasks:", tasks.TaskList)
	var taskList []*types.Tasks
	_ = copier.Copy(&taskList, tasks.TaskList)
	//logx.Error("api,taskList:", taskList)
	return &types.GetTasksResp{
		TasksList: taskList,
	}, nil
}
