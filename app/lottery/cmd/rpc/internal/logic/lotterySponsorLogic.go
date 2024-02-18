package logic

import (
	"context"
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
	lotteryId := in.Id
	// 1 根据lotteryId获取到uid和sponsorId
	uid, err := l.svcCtx.LotteryModel.FindUserIdByLotteryId(l.ctx, lotteryId)
	if err != nil {
		return nil, err
	}
	// 2 根据uid找到对应的user信息，在lotteryModel.go编写sql语句
	//  todo 根据uid和sponsorId从用户服务获取赞助商信息，返回
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
