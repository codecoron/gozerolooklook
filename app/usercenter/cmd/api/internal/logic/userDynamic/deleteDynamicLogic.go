package userDynamic

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteDynamicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteDynamicLogic {
	return &DeleteDynamicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteDynamicLogic) DeleteDynamic(req *types.DeleteDynamicReq) (resp *types.DeleteDynamicResp, err error) {
	// todo: add your logic here and delete this line

	return
}
