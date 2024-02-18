package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSponsorDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorDetailLogic {
	return &SponsorDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SponsorDetailLogic) SponsorDetail(in *pb.SponsorDetailReq) (*pb.SponsorDetailResp, error) {
	sponsorId := in.Id
	var res *model.UserSponsor
	res, err := l.svcCtx.UserSponsorModel.FindOne(l.ctx, sponsorId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "SponsorDetail, sponsorId:%v, error: %v", sponsorId, err)
	}
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR_NOT_FOUND), "SponsorDetail, sponsorId:%v, error: %v", sponsorId, err)
	}
	return &pb.SponsorDetailResp{
		Id:         res.Id,
		UserId:     res.UserId,
		Type:       res.Type,
		AppletType: res.AppletType,
		Name:       res.Name,
		Desc:       res.Desc,
		Avatar:     res.Avatar,
		IsShow:     res.IsShow,
		QrCode:     res.QrCode,
		InputA:     res.InputA,
		InputB:     res.InputB,
	}, nil
}
