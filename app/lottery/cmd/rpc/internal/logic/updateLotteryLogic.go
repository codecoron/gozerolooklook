package logic

import (
	"context"
	"github.com/pkg/errors"
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
	one, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询抽奖id失败 err: %v", err)
	}
	if one.UserId != in.UserId {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "该用户不是抽奖发起者")
	}
	lottery := new(model.Lottery)
	lottery.Id = in.Id
	lottery.PublishTime.Time = time.Now()
	lottery.PublishTime.Valid = true
	err = l.svcCtx.LotteryModel.UpdatePublishTime(l.ctx, lottery)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Lottery Database Exception lottery : %+v , err: %v", lottery, err)
	}

	return &pb.UpdateLotteryResp{}, nil
}
