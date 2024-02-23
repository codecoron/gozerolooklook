package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotterySponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLotterySponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotterySponsorLogic {
	return &LotterySponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LotterySponsorLogic) LotterySponsor(in *pb.LotterySponsorReq) (*pb.LotterySponsorResp, error) {
	detail, err := l.svcCtx.UserCenterRpc.SponsorDetail(l.ctx, &usercenter.SponsorDetailReq{
		Id: in.SponsorId,
	})
	if err != nil {
		return nil, err
	}
	sponsorInfo := new(pb.LotterySponsorResp)
	err = copier.Copy(sponsorInfo, detail)
	if err != nil {
		return nil, err
	}
	return sponsorInfo, nil
}
