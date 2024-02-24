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

	var CommentList []types.Comment
	if len(resp.Comment) > 0 {
		for _, item := range resp.Comment {
			var t types.Comment
			t.Id = item.Id
			t.UserId = item.UserId
			t.LotteryId = item.LotteryId
			t.PrizeName = item.PrizeName
			t.Content = item.Content
			t.Pics = item.Pics
			t.PraiseCount = item.PraiseCount
			CommentList = append(CommentList, t)
		}
	}

	return &types.CommentListResp{List: CommentList}, nil
}
