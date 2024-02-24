package comment

import (
	"context"
	"looklook/app/comment/cmd/rpc/comment"

	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCommentLogic) UpdateComment(req *types.CommentUpdateReq) (resp *types.CommentUpdateResp, err error) {
	_, err = l.svcCtx.CommentRpc.UpdateComment(l.ctx, &comment.UpdateCommentReq{
		Id:        req.Id,
		Content:   req.Content,
		Pics:      req.Pics,
		UserId:    req.UserId,
		LotteryId: req.LotteryId,
		PrizeName: req.PrizeName,
	})
	if err != nil {
		return nil, err
	}

	return &types.CommentUpdateResp{}, nil
}
