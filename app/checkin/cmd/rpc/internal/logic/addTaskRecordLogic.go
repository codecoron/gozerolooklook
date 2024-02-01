package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"looklook/app/checkin/model"
	"looklook/common/xerr"

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
	_, err := l.svcCtx.TasksModel.FindOne(l.ctx, in.TaskId)
	if err == sqlc.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_NOT_FOUND), "任务不存在")
	} else if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询任务内容失败")
	}
	taskRecord := new(model.TaskRecord)
	taskRecord.TaskId = in.TaskId
	taskRecord.UserId = in.UserId
	taskRecord.Type = in.Type
	taskRecord.IsFinished = 1
	_, err = l.svcCtx.TaskRecordModel.Insert(l.ctx, taskRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "新增任务完成记录失败")
	}
	return &pb.AddTaskRecordResp{}, nil
}
