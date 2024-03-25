package user

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) DetailLogic {
	return DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req types.UserInfoReq) (*types.UserInfoResp, error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)

	userInfoResp, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
		Id: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("查询用户详情失败"), "查询用户详情失败 err : %v , userId : %d  , userInfoResp : %+v", err, userId, userInfoResp)
	}
	userLotteryInfoResp, err := l.svcCtx.LotteryRpc.GetLotteryStatistic(l.ctx, &lottery.GetLotteryStatisticReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("查询用户抽奖详情失败"), "查询用户抽奖详情失败 err : %v , userId : %d  , userLotteryInfoResp : %+v", err, userId, userLotteryInfoResp)
	}
	logx.Error("userLotteryInfoResp: ", userLotteryInfoResp.CreatedCount, userLotteryInfoResp.WonCount, userLotteryInfoResp.ParticipationCount)

	userCheckinInfoIntegralResp, err := l.svcCtx.CheckinRpc.GetIntegralByUserId(l.ctx, &checkin.GetIntegralByUserIdReq{
		UserId: userId,
	})
	logx.Error("userCheckinInfoIntegralResp: ", userCheckinInfoIntegralResp)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("查询用户心愿详情失败"), "查询用户心愿详情失败 err : %v , userId : %d  , userCheckinInfoIntegralResp : %+v", err, userId, userCheckinInfoIntegralResp)
	}

	var userInfo types.User
	_ = copier.Copy(&userInfo, userInfoResp.User)
	_ = copier.Copy(&userInfo, userLotteryInfoResp)
	_ = copier.Copy(&userInfo, userCheckinInfoIntegralResp)

	return &types.UserInfoResp{
		UserInfo: userInfo,
	}, nil
}
