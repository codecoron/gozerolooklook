package logic

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/model"
	"looklook/common/xerr"
	"time"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLotteryLogic {
	return &UpdateLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLotteryLogic) UpdateLottery(in *pb.UpdateLotteryReq) (*pb.UpdateLotteryResp, error) {
	err := l.svcCtx.LotteryModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		one, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
		if err != nil {
			logx.Error("查询抽奖id失败:%v", err)
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err: %v", err)
		}
		if one.UserId != in.UserId {
			logx.Error("user_id与抽奖id不匹配")
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user_id与抽奖id不匹配")
		}
		lottery := new(model.Lottery)
		lottery.Id = in.Id
		var pTime sql.NullTime
		var aTime sql.NullTime
		if in.PublishTime != 0 {
			pTime.Time = time.Unix(in.PublishTime, 0)
			pTime.Valid = true
		} else {
			pTime.Valid = false
		}
		if in.PublishTime != 0 {
			aTime.Time = time.Unix(in.AwardDeadline, 0)
			aTime.Valid = true
		} else {
			aTime.Valid = false
		}
		lottery.PublishTime = pTime
		lottery.AwardDeadline = aTime

		_, err = l.svcCtx.LotteryModel.TransUpdate(l.ctx, session, lottery)
		if err != nil {
			logx.Error("修改失败:%v", err)
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Lottery Database Exception lottery : %+v , err: %v", lottery, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateLotteryResp{}, nil
}
