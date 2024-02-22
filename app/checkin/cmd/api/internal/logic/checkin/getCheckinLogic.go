package checkin

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/common/ctxdata"
)

type GetCheckinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCheckinLogic {
	return &GetCheckinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCheckinLogic) GetCheckin(req *types.GetCheckinReq) (resp *types.GetCheckinResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	record, err := l.svcCtx.CheckinRpc.GetCheckinRecordByUserId(l.ctx, &checkin.GetCheckinRecordByUserIdReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetCheckinResp{
		ContinuousCheckinDays: record.ContinuousCheckinDays,
		State:                 record.State,
		Integral:              record.Integral,
		SubStatus:             record.SubStatus,
	}, nil
}
