package userSponsor

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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
	logx.Debug(sponsor)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("修改联系方式失败"), "add Sponsor rpc AddUserSponsor fail req: %+v , err : %v ", req, err)
	}
	return &types.UpdateSponsorResp{}, nil
	return
}
