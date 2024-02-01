package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/checkin/model"
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
	checkinRecord, err := l.svcCtx.CheckinRecordModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find check-in data : %+v , err: %v", checkinRecord, err)
	}
	integarl, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find integral data : %+v , err: %v", integarl, err)
	}
	err = l.svcCtx.CheckinRecordModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		if checkinRecord.State == 1 {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "今日已签到")
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
		checkinRecord.LastCheckinDate.Valid = true
		checkinRecord.State = 1
		err = l.svcCtx.CheckinRecordModel.TransUpdateByUserId(l.ctx, session, checkinRecord)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update check-in data : %+v , err: %v", checkinRecord, err)
		}
		err = l.svcCtx.IntegralModel.TransUpdateByUserId(l.ctx, session, integarl)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update integral data : %+v , err: %v", integarl, err)
		}

		// 增加心愿值增加记录
		integralRecord := new(model.IntegralRecord)
		integralRecord.Integral = i
		integralRecord.Content = "签到"
		integralRecord.UserId = in.UserId
		_, err = l.svcCtx.IntegralRecordModel.TransInsert(l.ctx, session, integralRecord)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to insert integralRecord data : %+v , err: %v", integralRecord, err)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &pb.UpdateCheckinRecordResp{
		State:                 checkinRecord.State,
		ContinuousCheckinDays: checkinRecord.ContinuousCheckinDays,
		Integral:              integarl.Integral,
	}, nil
}
