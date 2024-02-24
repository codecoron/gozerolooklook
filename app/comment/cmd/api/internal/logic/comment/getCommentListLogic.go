package comment

import (
	"context"
	"looklook/app/comment/cmd/rpc/comment"

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
		// todo 获取用户信息s
		//userInfo, err := l.svcCtx.UsercenterRpc.GetUserInfoList(l.ctx, &comment.GetUserInfoListReq{})
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
	}

	return &types.CommentListResp{List: CommentList}, nil
}
