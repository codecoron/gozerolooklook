package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchParticipationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchParticipationLogic {
	return &SearchParticipationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchParticipationLogic) SearchParticipation(req *types.SearchLotteryParticipationReq) (resp *types.SearchLotteryParticipationResp, err error) {
	r, err := l.svcCtx.LotteryRpc.SearchLotteryParticipation(l.ctx, &lottery.SearchLotteryParticipationReq{
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, err
	}

	userInfos := []*usercenter.User{}
	for i := range r.List {
		userId := r.List[i].UserId
		info, err := l.svcCtx.UsercenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
			Id: userId,
		})
		if err != nil {
			return nil, err
		}
		userInfos = append(userInfos, info.User)
	}
	err = copier.Copy(&resp.List, userInfos)
	return
}
