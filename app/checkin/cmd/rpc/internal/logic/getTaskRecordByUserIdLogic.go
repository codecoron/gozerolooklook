package logic

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jinzhu/copier"
	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskRecordByUserIdLogic {
	return &GetTaskRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskRecordByUserIdLogic) GetTaskRecordByUserId(in *pb.GetTaskRecordByUserIdReq) (*pb.GetTaskRecordByUserIdResp, error) {
	// 查询所有任务列表
	query := squirrel.Select().From("tasks")
	tasks, err := l.svcCtx.TasksModel.FindAll(l.ctx, query, "id ASC")
	if err != nil {
		return nil, err
	}
	var taskList []*pb.Tasks
	_ = copier.Copy(&taskList, tasks)

	query = squirrel.Select().From("task_record")
	finishTasks, err := l.svcCtx.TaskRecordModel.FindByUserId(l.ctx, in.UserId, query, "id ASC")

	// 赋值任务的完成情况
	fMap := make(map[int64]int64)
	for _, fTask := range finishTasks {
		fMap[fTask.TaskId] = fTask.IsFinished
	}
	for i := range taskList {
		if isFinished, ok := fMap[taskList[i].Id]; ok {
			taskList[i].IsFinished = isFinished
		}
		taskList[i].Count = -1
		taskList[i].NeedCount = -1
	}
	return &pb.GetTaskRecordByUserIdResp{
		TaskList: taskList,
	}, nil
}
