package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"
	"looklook/app/checkin/model"
	"looklook/common/xerr"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCheckinRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCheckinRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCheckinRecordByUserIdLogic {
	return &GetCheckinRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCheckinRecordByUserIdLogic) GetCheckinRecordByUserId(in *pb.GetCheckinRecordByUserIdReq) (*pb.GetCheckinRecordByUserIdResp, error) {
	checkinRecord := new(model.CheckinRecord)
	integarl := new(model.Integral)

	err := l.svcCtx.CheckinRecordModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		// 根据user_id查询用户的签到记录，如果没有就创建
		getCheckinRecord, err := l.svcCtx.CheckinRecordModel.FindOneByUserId(l.ctx, in.UserId)
		if err == sqlc.ErrNotFound { // 没查询到用户的数据，说明用户从来没有签到过，新增记录
			// 新增签到记录
			checkinRecord.UserId = in.UserId
			insert, err := l.svcCtx.CheckinRecordModel.TransInsertByUserId(l.ctx, session, checkinRecord)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to insert check-in data : %+v , err: %v", insert, err)
			}
			// 新增积分记录
			integarl.UserId = in.UserId
			insert, err = l.svcCtx.IntegralModel.TransInsertByUserId(l.ctx, session, integarl)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to insert integarl data : %+v , err: %v", insert, err)
			}
			// todo:新增用户任务进度记录
			return nil
		} else if err != nil {
			// 其他错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find check-in data : %+v , err: %v", getCheckinRecord, err)
		}

		// 查询积分，任务列表
		getIntegral, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
		logx.Error("getIntegral:", getIntegral)
		// 将getCheckinRecord 复制到 checkinRecord
		_ = copier.Copy(checkinRecord, getCheckinRecord)
		_ = copier.Copy(integarl, getIntegral)

		// 将现在的时间转换为UTC时间，然后截断为当天的起始时间，只需要知道日期就行
		today := time.Now().UTC().Truncate(24 * time.Hour)
		//logx.Error("现在的日期today:", today)
		// 也是截断为当天的起始时间
		targetDate := getCheckinRecord.LastCheckinDate.Time.Truncate(24 * time.Hour)
		//logx.Error("上次签到的日期targetDate:", targetDate)

		switch {
		case targetDate.Equal(today):
			// 如果今天签了到，什么都不用变，因为签到的时候会更新
			return nil
		case targetDate.Equal(today.Add(-24 * time.Hour)):
			// 如果是昨天签的到，然后刚好是第七天，天数归零
			if checkinRecord.ContinuousCheckinDays >= 7 {
				checkinRecord.ContinuousCheckinDays = 0
				checkinRecord.State = 0
				err := l.svcCtx.CheckinRecordModel.TransUpdateByUserId(l.ctx, session, checkinRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update check-in data : %+v , err: %v", checkinRecord, err)
				}
			} else if checkinRecord.State != 0 { // 如果是昨天签的到，state还是1的话，变为0
				checkinRecord.State = 0
				err := l.svcCtx.CheckinRecordModel.TransUpdateByUserId(l.ctx, session, checkinRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update check-in data : %+v , err: %v", checkinRecord, err)
				}
			}
			return nil
		default:
			if checkinRecord.State != 0 || checkinRecord.ContinuousCheckinDays != 0 {
				checkinRecord.ContinuousCheckinDays = 0
				checkinRecord.State = 0
				err := l.svcCtx.CheckinRecordModel.TransUpdateByUserId(l.ctx, session, checkinRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update check-in data : %+v , err: %v", checkinRecord, err)
				}
			}
			return nil
		}
	})

	if err != nil {
		return nil, err
	}
	logx.Error("integarl.Integral: ", integarl.Integral)
	return &pb.GetCheckinRecordByUserIdResp{
		ContinuousCheckinDays: checkinRecord.ContinuousCheckinDays,
		State:                 checkinRecord.State,
		Integral:              integarl.Integral,
	}, nil
}
