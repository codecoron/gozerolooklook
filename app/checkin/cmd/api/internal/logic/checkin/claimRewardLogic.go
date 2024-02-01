package checkin

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClaimRewardLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClaimRewardLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClaimRewardLogic {
	return &ClaimRewardLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClaimRewardLogic) ClaimReward(req *types.ClaimRewardReq) (resp *types.ClaimRewardResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.CheckinRpc.UpdateTaskRecord(l.ctx, &checkin.UpdateTaskRecordReq{
		UserId: userId,
		TaskId: req.TaskId,
	})

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("claim reward fail"), "claim reward rpc ClaimReward fail req: %+v , err : %v ", req, err)
	}

	return
}
