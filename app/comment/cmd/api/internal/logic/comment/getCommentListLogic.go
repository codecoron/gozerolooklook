package comment

import (
	"context"
	"looklook/app/comment/cmd/rpc/comment"
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
		for _, item := range resp.Comment {
			// 先得到所有用户Id的切片，传入这个切片得到用户信息列表
			userIds = append(userIds, item.UserId)
		}
		// 根据用户Id获取用户信息
		var userInfoList []types.User
		userInfoList = make([]types.User, len(userIds))
		// todo 获取用户信息s： 传入用户Id的切片，得到用户信息列表
		/// 1. 传入用户Id的切片，得到用户信息列表

		//userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfoList(l.ctx, &comment.GetUserInfoListReq{})
		// 捏造用户信息
		for idx, item := range userInfoList {
			item.Id = userIds[idx]
			item.Nickname = "test"
			item.Avatar = "https://jhonronxton-lottery.oss-cn-beijing.aliyuncs.com/d08dd5ea-b660-4020-b7bc-82e219a314b4.jpg"
			userInfoList[idx] = item
		}
		// 名字打码，只留下字符的第一个和最后一个，中间多个字符只有两个*
		for idx, item := range userInfoList {
			if len(item.Nickname) > 2 {
				item.Nickname = item.Nickname[:1] + "**" + item.Nickname[len(item.Nickname)-1:]
			} else {
				item.Nickname = item.Nickname[:1] + "**"
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
