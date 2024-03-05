package userSponsor

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/pb"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSponsorDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDelLogic {
	return &SponsorDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorDelLogic) SponsorDel(req *types.SponsorDelReq) (resp *types.SponsorDelResp, err error) {
	// todo: add your logic here and delete this line

	_, err = l.svcCtx.UsercenterRpc.DelUserSponsor(l.ctx, &pb.DelUserSponsorReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
