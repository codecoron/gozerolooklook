package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/common/xerr"
	"time"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCheckinRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCheckinRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCheckinRecordLogic {
	return &UpdateCheckinRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCheckinRecordLogic) UpdateCheckinRecord(in *pb.UpdateCheckinRecordReq) (*pb.UpdateCheckinRecordResp, error) {
	// todo: add your logic here and delete this line
	checkinRecord, err := l.svcCtx.CheckinRecordModel.FindOneByUserId(l.ctx, in.UserId)
	integarl, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
	if checkinRecord.State == 1 {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "今日已签到")
	}
	var i int64
	switch checkinRecord.ContinuousCheckinDays {
	case 0, 1:
		i = 5
	case 2:
		i = 10
	case 3:
		i = 15
	case 4:
		i = 20
	case 5:
		i = 30
	case 6:
		i = 40
	default:
		// 如果 `checkinRecord.ContinuousCheckinDays` 的值不在上述范围内，则设置默认值
		i = 0
	}
	integarl.Integral += i
	checkinRecord.ContinuousCheckinDays += 1
	checkinRecord.LastCheckinDate.Time = time.Now()
	checkinRecord.State = 1
	err = l.svcCtx.CheckinRecordModel.Update(l.ctx, checkinRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update check-in data : %+v , err: %v", checkinRecord, err)
	}
	err = l.svcCtx.IntegralModel.Update(l.ctx, integarl)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update integral data : %+v , err: %v", checkinRecord, err)
	}
	return &pb.UpdateCheckinRecordResp{
		State:                 checkinRecord.State,
		ContinuousCheckinDays: checkinRecord.ContinuousCheckinDays,
		Integral:              integarl.Integral,
	}, nil
}
