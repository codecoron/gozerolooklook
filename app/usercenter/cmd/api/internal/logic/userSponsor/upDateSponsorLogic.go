package userSponsor

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
	"looklook/app/usercenter/cmd/rpc/pb"
)

type UpDateSponsorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateSponsorLogic {
	return &UpDateSponsorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpDateSponsorLogic) UpDateSponsor(req *types.UpdateSponsorReq) (resp *types.UpdateSponsorResp, err error) {
	// todo: add your logic here and delete this line
	pbSponsorReq := new(pb.UpdateUserSponsorReq)
	err = copier.Copy(pbSponsorReq, req)
	if err != nil {
		return nil, err
	}
	sponsor, err := l.svcCtx.UsercenterRpc.UpdateUserSponsor(l.ctx, pbSponsorReq)
	if err != nil {
		return nil, errors.Wrapf(err, "update Sponsor rpc UpdateUserSponsor fail req: %+v , err : %v ", req, err)
	}

	resp = &types.UpdateSponsorResp{}
	err = copier.Copy(sponsor, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
