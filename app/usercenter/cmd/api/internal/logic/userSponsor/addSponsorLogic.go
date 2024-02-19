package userSponsor

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/ctxdata"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSponsorLogic {
	return &AddSponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSponsorLogic) AddSponsor(req *types.CreateSponsorReq) (resp *types.CreateSponsorResp, err error) {
	pbSponsor := new(pb.AddUserSponsorReq)
	err = copier.Copy(pbSponsor, req)
	if err != nil {
		return nil, err
	}
	pbSponsor.UserId = ctxdata.GetUidFromCtx(l.ctx)
	sponsor, err := l.svcCtx.UsercenterRpc.AddUserSponsor(l.ctx, pbSponsor)
	if err != nil {
		return nil, err
	}
	return &types.CreateSponsorResp{
		Id: sponsor.Id,
	}, nil
}
