package userSponsor

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSponsorDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDetailLogic {
	return &SponsorDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorDetailLogic) SponsorDetail(req *types.SponosorDetailReq) (resp *types.SponosorDetailResp, err error) {
	sponsorResp := new(pb.SponsorDetailResp)
	sponsorResp, err = l.svcCtx.UsercenterRpc.SponsorDetail(l.ctx, &usercenter.SponsorDetailReq{
		Id: req.Id,
	})
	resp = &types.SponosorDetailResp{
		Id:         sponsorResp.Id,
		UserId:     sponsorResp.UserId,
		Type:       sponsorResp.Type,
		AppletType: sponsorResp.AppletType,
		IsShow:     sponsorResp.IsShow,
		Name:       sponsorResp.Name,
		Desc:       sponsorResp.Desc,
		Avatar:     sponsorResp.Avatar,
		QrCode:     sponsorResp.QrCode,
		InputA:     sponsorResp.InputA,
		InputB:     sponsorResp.InputB,
	}
	return resp, nil
}
