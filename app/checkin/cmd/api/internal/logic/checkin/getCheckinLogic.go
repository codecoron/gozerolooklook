package checkin

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, errors.Wrapf(xerr.NewErrMsg("get checkin fail"), "get checkin rpc GetCheckin fail req: %+v , err : %v ", req, err)
	}

	return &types.GetCheckinResp{
		ContinuousCheckinDays: record.ContinuousCheckinDays,
		State:                 record.State,
		Integral:              record.Integral,
	}, nil
}
