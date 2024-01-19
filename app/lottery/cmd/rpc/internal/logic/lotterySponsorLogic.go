package logic

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

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
	// todo: add your logic here and delete this line
	lotteryId := in.Id
	// 1 根据lotteryId获取到uid
	uid, err := l.svcCtx.LotteryModel.FindUserIdByLotteryId(l.ctx, lotteryId)
	if err != nil {
		return nil, err
	}
	// 2 根据uid找到对应的user信息，在lotteryModel.go编写sql语句
	userInfo, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: *uid,
	})
	if err != nil {
		return nil, err
	}
	sponsorInfo := new(pb.LotterySponsorResp)
	sponsorInfo.Id = userInfo.User.Id
	sponsorInfo.NickName = userInfo.User.Nickname
	sponsorInfo.Avatar = userInfo.User.Avatar
	sponsorInfo.Info = userInfo.User.Info
	return sponsorInfo, nil
}
