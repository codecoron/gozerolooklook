package userContact

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/pb"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContactDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContactDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContactDelLogic {
	return &ContactDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContactDelLogic) ContactDel(req *types.ContactDelReq) (resp *types.ContactDelResp, err error) {
	_, err = l.svcCtx.UsercenterRpc.DelUserContact(l.ctx, &pb.DelUserContactReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
