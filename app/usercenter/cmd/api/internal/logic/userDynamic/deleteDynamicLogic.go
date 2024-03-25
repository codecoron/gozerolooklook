package userDynamic

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/pb"

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
	_, err = l.svcCtx.UsercenterRpc.DelUserDynamic(l.ctx, &pb.DelUserDynamicReq{
		Id:     req.Id,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return
}
