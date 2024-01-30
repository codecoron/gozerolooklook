package checkin

import (
	"context"

	"looklook/app/checkin/cmd/api/internal/svc"
	"looklook/app/checkin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FinishTaskLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFinishTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FinishTaskLogic {
	return &FinishTaskLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FinishTaskLogic) FinishTask(req *types.FinishTaskReq) (resp *types.FinishTaskResp, err error) {
	// todo: add your logic here and delete this line

	return
}
