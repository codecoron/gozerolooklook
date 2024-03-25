package lottery

import (
	"context"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateClockTaskRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateClockTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateClockTaskRecordLogic {
	return &CreateClockTaskRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateClockTaskRecordLogic) CreateClockTaskRecord(req *types.CreateClockTaskRecordReq) (resp *types.CreateClockTaskRecordResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	addClockTaskRecord, err := l.svcCtx.LotteryRpc.AddClockTaskRecord(l.ctx, &lottery.AddClockTaskRecordReq{
		UserId:      userId,
		LotteryId:   req.LotteryId,
		ClockTaskId: req.ClockTaskId,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateClockTaskRecordResp{
		Id: addClockTaskRecord.Id,
	}, nil
}
