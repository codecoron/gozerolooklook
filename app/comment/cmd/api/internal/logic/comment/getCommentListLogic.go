package comment

import (
	"context"
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentListLogic {
	return &GetCommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentListLogic) GetCommentList(req *types.CommentListReq) (*types.CommentListResp, error) {
	resp, err := l.svcCtx.CommentRpc.SearchComment(l.ctx, &comment.SearchCommentReq{
		LastId: req.LastId,
		Page:   req.Page,
		Limit:  req.PageSize,
		Sort:   req.Sort,
	})
	if err != nil {
		return nil, err
	}

	var CommentList []types.Comments

	if len(resp.Comment) > 0 {
		// 获取到所有用户Id，根据用户Id获取用户信息
		userIds := make([]int64, 0)
		// 根据用户Id获取用户信息
		userInfoList := make([]types.User, len(userIds))
		// 用map存储用户信息，去重
		userInfoMap := make(map[int64]*types.User)
		for _, item := range resp.Comment {
			// 先得到所有用户Id的切片，传入这个切片得到用户信息列表
			userInfoList = append(userInfoList, types.User{
				Id: item.UserId,
			})

			// 去重
			if _, ok := userInfoMap[item.UserId]; !ok {
				userInfoMap[item.UserId] = &types.User{
					Id: item.UserId,
				}
				userIds = append(userIds, item.UserId)
			}
		}

		/// 1. 传入用户Id的切片，得到用户信息列表

		userInfos, err := l.svcCtx.UsercenterRpc.GetUserInfoByUserIds(l.ctx, &usercenter.GetUserInfoByUserIdsReq{
			UserIds: userIds,
		})
		if err != nil {
			return nil, err
		}
		// 2. 将用户信息列表转换为map
		for _, item := range userInfos.UserInfo {
			userInfoMap[item.Id] = &types.User{
				Id:       item.Id,
				Nickname: item.Nickname,
				Avatar:   item.Avatar,
			}
		}
		// 3. 将用户信息列表转换为切片
		for idx, _ := range userInfoList {
			userId := userInfoList[idx].Id
			userInfoList[idx] = *userInfoMap[userId]
		}

		// 打印测试
		//fmt.Println("userInfoList:", userInfoList)

		// 名字打码，只留下字符的第一个和最后一个，中间多个字符只有两个*
		for idx, item := range userInfoList {
			if len(item.Nickname) > 2 {
				item.Nickname = item.Nickname[:1] + "**" + item.Nickname[len(item.Nickname)-1:]
			} else {
				item.Nickname = item.Nickname[:] + "**"
			}
			userInfoList[idx] = item
		}
		for idx, item := range resp.Comment {
			var t types.Comments
			t.Id = item.Id
			t.UserId = item.UserId
			t.LotteryId = item.LotteryId
			t.PrizeName = item.PrizeName
			t.Content = item.Content
			t.Pics = item.Pics
			t.PraiseCount = item.PraiseCount
			t.CreateTime = item.CreateTime
			t.UpdateTime = item.UpdateTime
			t.User = userInfoList[idx]
			CommentList = append(CommentList, t)
		}

		// 得到CommentIds，根据CommentIds获取点赞信息
		CommentIds := make([]int64, 0)
		for _, item := range resp.Comment {
			CommentIds = append(CommentIds, item.Id)
		}
		userId := ctxdata.GetUidFromCtx(l.ctx)
		list, err := l.svcCtx.CommentRpc.IsPraiseList(l.ctx, &comment.IsPraiseListReq{
			CommentId: CommentIds,
			UserId:    userId,
		})
		if err != nil {
			return nil, err
		}
		// 将list.PraiseId转换为map
		praiseMap := make(map[int64]int)
		for _, v := range list.PraiseId {
			praiseMap[v] = 1
		}
		for idx, item := range CommentList {
			item.IsPraise = 0
			if _, ok := praiseMap[item.Id]; ok {
				item.IsPraise = 1
			}
			CommentList[idx] = item
		}
	}

	return &types.CommentListResp{List: CommentList}, nil
}
