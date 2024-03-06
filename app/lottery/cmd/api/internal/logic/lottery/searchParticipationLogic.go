package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/usercenter/cmd/rpc/usercenter"
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
		PageIndex: req.PageIndex,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	var userIds []int64

	for i := range r.List {
		userIds = append(userIds, r.List[i].UserId)
	}

	userInfos := new(usercenter.GetUserInfoByUserIdsResp)
	if len(userIds) > 0 {
		userInfos, err = l.svcCtx.UsercenterRpc.GetUserInfoByUserIds(l.ctx, &usercenter.GetUserInfoByUserIdsReq{
			UserIds: userIds,
		})
		if err != nil {
			return nil, err
		}
	}

	// 名字打码，只留下字符的第一个和最后一个，中间多个字符只有两个*
	for idx, item := range userInfos.UserInfo {
		if len(item.Nickname) > 2 {
			item.Nickname = item.Nickname[:1] + "**" + item.Nickname[len(item.Nickname)-1:]
		} else {
			item.Nickname = item.Nickname[:] + "**"
		}
		userInfos.UserInfo[idx] = item
	}

	resp = new(types.SearchLotteryParticipationResp)
	err = copier.Copy(&resp.List, userInfos.UserInfo)
	if err != nil {
		return nil, err
	}

	resp.Count = r.Count
	return
}
