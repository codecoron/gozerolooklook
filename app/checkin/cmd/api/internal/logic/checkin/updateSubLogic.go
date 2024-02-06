package checkin

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/common/ctxdata"

	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSubLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSubLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSubLogic {
	return &UpdateSubLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSubLogic) UpdateSub(req *types.UpdateSubReq) (resp *types.UpdateSubResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.CheckinRpc.UpdateSub(l.ctx, &checkin.UpdateSubReq{
		UserId: userId,
		State:  req.State,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return
}
