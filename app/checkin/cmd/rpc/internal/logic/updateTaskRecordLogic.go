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
	task, err := l.svcCtx.TasksModel.FindOne(l.ctx, in.TaskId)
	if err == sqlc.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_NOT_FOUND), "任务不存在")
	} else if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询任务失败")
	}

	taskRecord, err := l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(l.ctx, in.UserId, in.TaskId)
	if err == sqlc.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_NOT_FINISHED), "任务未完成，taskRecord : %v", taskRecord)
	} else if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询任务记录失败，taskRecord : %v", taskRecord)
	}
	if taskRecord.IsFinished == 2 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_TASK_REWARD_COLLECTED), "不可重复领取奖励")
	}
	taskRecord.IsFinished = 2
	err = l.svcCtx.TaskRecordModel.Update(l.ctx, taskRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "领取任务奖励失败")
	}

	// 增加心愿值增加记录
	integralRecord := new(model.IntegralRecord)
	integralRecord.Integral = task.Integral
	integralRecord.Content = "任务奖励"
	integralRecord.UserId = in.UserId
	_, err = l.svcCtx.IntegralRecordModel.Insert(l.ctx, integralRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to insert integralRecord data : %+v , err: %v", integralRecord, err)
	}

	// 增加用户心愿值
	integarl, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "integral : %+v , err: %v", integarl, err)
	}
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_RECORD_NOT_FOUND), "integral NOT FOUND: %+v , err: %v", integarl, err)
	}
	integarl.Integral += task.Integral
	err = l.svcCtx.IntegralModel.UpdateByUserId(l.ctx, integarl)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update integral data : %+v , err: %v", integarl, err)
	}
	return &pb.UpdateTaskRecordResp{}, nil
}
