package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"looklook/common/xerr"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubLogic {
	return &UpdateSubLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSubLogic) UpdateSub(in *pb.UpdateSubReq) (*pb.UpdateSubResp, error) {
	getProgress, err := l.svcCtx.TaskProgressModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskProgress data : %+v , err: %v", getProgress, err)
	}
	getProgress.IsSubCheckin = in.State
	err = l.svcCtx.TaskProgressModel.UpdateByUserId(l.ctx, getProgress)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update taskProgress data : %+v , err: %v", getProgress, err)
	}
	// 完成任务2
	if in.State == 1 {
		// 查询任务记录，如果有说明已完成，没有就添加记录
		_, err = l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(l.ctx, in.UserId, 2)
		if err == sqlc.ErrNotFound {
			addTaskRecord := &pb.AddTaskRecordReq{
				UserId: in.UserId,
				TaskId: 2,
			}
			logic := NewAddTaskRecordLogic(l.ctx, l.svcCtx)
			_, err := logic.AddTaskRecord(addTaskRecord)
			if err != nil {
				return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 2, err: %v", err)
			}
		}
	}
	return &pb.UpdateSubResp{}, nil
}
