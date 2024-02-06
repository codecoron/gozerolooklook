package checkin

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
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
	userId := ctxdata.GetUidFromCtx(l.ctx)
	// 查询用户任务进度，返回具体数量
	count, err := l.svcCtx.CheckinRpc.GetTaskProgress(l.ctx, &checkin.GetTaskProgressReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	// 查询用户完成的任务
	tasks, err := l.svcCtx.CheckinRpc.GetTaskRecordByUserId(l.ctx, &checkin.GetTaskRecordByUserIdReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	//logx.Error("api,tasks:", tasks.TaskList)
	var taskList []*types.Tasks
	_ = copier.Copy(&taskList, tasks.TaskList)
	//logx.Error("api,taskList:", taskList)
	// 返回任务进度具体数量
	return &types.GetTasksResp{
		TasksList: taskList,
		DayCount:  count.DayCount,
		WeekCount: count.WeekCount,
	}, nil
}
