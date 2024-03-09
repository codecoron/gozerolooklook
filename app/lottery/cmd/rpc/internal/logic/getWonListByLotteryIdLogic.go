package logic

import (
	"context"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"sort"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWonListByLotteryIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWonListByLotteryIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWonListByLotteryIdLogic {
	return &GetWonListByLotteryIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type UserInfo struct {
	Id       int64
	Nickname string
	Avatar   string
}

func (l *GetWonListByLotteryIdLogic) GetWonListByLotteryId(in *pb.GetWonListByLotteryIdReq) (*pb.GetWonListByLotteryIdResp, error) {
	// 查询当前抽奖的中奖列表
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("lottery_id = ?", in.LotteryId)
	WinInfo, err := l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}

	// 统计userIds
	userIds := make([]int64, 0)
	for _, v := range WinInfo {
		userIds = append(userIds, v.UserId)
	}

	// 通过Wininfo统计奖品Id下的中奖者id，存入map
	prizeUserMap := make(map[int64][]int64)
	for _, v := range WinInfo {
		if _, ok := prizeUserMap[v.PrizeId]; !ok {
			prizeUserMap[v.PrizeId] = make([]int64, 0)
		}
		prizeUserMap[v.PrizeId] = append(prizeUserMap[v.PrizeId], v.UserId)
	}

	// 查询用户信息
	if len(userIds) == 0 {
		return &pb.GetWonListByLotteryIdResp{}, nil
	}
	userInfo, err := l.svcCtx.UserCenterRpc.GetUserInfoByUserIds(l.ctx, &usercenter.GetUserInfoByUserIdsReq{UserIds: userIds})
	if err != nil {
		return nil, err
	}

	// 名字打码，只留下字符的第一个和最后一个，中间多个字符只有两个*
	for idx, item := range userInfo.UserInfo {
		if len(item.Nickname) > 2 {
			item.Nickname = item.Nickname[:1] + "**" + item.Nickname[len(item.Nickname)-1:]
		} else {
			item.Nickname = item.Nickname[:] + "**"
		}
		userInfo.UserInfo[idx] = item
	}

	// 转为map
	userInfoMap := make(map[int64]*UserInfo)
	for _, v := range userInfo.UserInfo {
		userInfoMap[v.Id] = &UserInfo{
			Id:       v.Id,
			Nickname: v.Nickname,
			Avatar:   v.Avatar,
		}
	}

	// 根据抽奖id查询奖品信息
	prizeInfo, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, in.LotteryId)
	if err != nil {
		return nil, err
	}

	// prizeInfo 按照 level 从小到大排序
	// 1. 遍历prizeInfo，将prizeInfo按照level排序，从小到大排序
	sort.Slice(prizeInfo, func(i, j int) bool {
		return prizeInfo[i].Level < prizeInfo[j].Level
	})
	// 2. 将排序后的prizeInfo放入list
	list := make([]*pb.WonList2, 0)
	for _, v := range prizeInfo {
		prize := &pb.Prize{
			Id:        v.Id,
			LotteryId: v.LotteryId,
			Type:      v.Type,
			Name:      v.Name,
			Level:     v.Level,
			Thumb:     v.Thumb,
			Count:     v.Count,
			GrantType: v.GrantType,
		}
		users := make([]*pb.UserInfo, 0)
		for _, userId := range prizeUserMap[v.Id] {
			user := userInfoMap[userId]
			users = append(users, &pb.UserInfo{
				Id:       user.Id,
				Nickname: user.Nickname,
				Avatar:   user.Avatar,
			})
		}
		list = append(list, &pb.WonList2{
			Prize:       prize,
			WinnerCount: int64(len(prizeUserMap[v.Id])),
			Users:       users,
		})
	}

	return &pb.GetWonListByLotteryIdResp{
		List: list,
	}, nil
}
