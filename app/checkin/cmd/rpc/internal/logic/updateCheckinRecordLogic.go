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
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "checkinRecord : %+v , err: %v", checkinRecord, err)
	}
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_RECORD_NOT_FOUND), "checkinRecord NOT FOUND: %+v , err: %v", checkinRecord, err)
	}
	integarl, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "integral : %+v , err: %v", integarl, err)
	}
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_RECORD_NOT_FOUND), "integral NOT FOUND: %+v , err: %v", integarl, err)
	}
	err = l.svcCtx.CheckinRecordModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		if checkinRecord.State == 1 {
			return errors.Wrapf(xerr.NewErrCode(xerr.CHECKIN_REPEAT), "err : %v", err)
		}
		// 更新积分值
		integrals := calculateCheckinIntegral(checkinRecord.ContinuousCheckinDays)
		integarl.Integral += integrals
		err = l.svcCtx.IntegralModel.TransUpdateByUserId(l.ctx, session, integarl)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update integral data : %+v , err: %v", integarl, err)
		}
		// 更新签到状态
		checkinRecord.ContinuousCheckinDays += 1
		checkinRecord.LastCheckinDate.Time = time.Now()
		checkinRecord.LastCheckinDate.Valid = true
		checkinRecord.State = 1
		err = l.svcCtx.CheckinRecordModel.TransUpdateByUserId(l.ctx, session, checkinRecord)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update check-in data : %+v , err: %v", checkinRecord, err)
		}
		// 添加心愿值增减记录
		integralRecord := new(model.IntegralRecord)
		integralRecord.Integral = integrals
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

func calculateCheckinIntegral(continuousCheckinDays int64) int64 {
	var integral int64
	switch continuousCheckinDays {
	case 0, 1:
		integral = 5
	case 2:
		integral = 10
	case 3:
		integral = 15
	case 4:
		integral = 20
	case 5:
		integral = 30
	case 6:
		integral = 40
	default:
		integral = 0
	}
	return integral
}
