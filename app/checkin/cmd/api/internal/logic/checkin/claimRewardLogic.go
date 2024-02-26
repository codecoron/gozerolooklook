package checkin

import (
	"context"
	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/common/ctxdata"

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
	// todo: 错误处理
	if err != nil {
		return nil, err
	}

	return
}
