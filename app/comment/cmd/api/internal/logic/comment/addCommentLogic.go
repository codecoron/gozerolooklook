package comment

import (
	"context"
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/common/ctxdata"

	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentLogic) AddComment(req *types.CommentAddReq) (resp *types.CommentAddResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.CommentRpc.AddComment(l.ctx, &comment.AddCommentReq{
		UserId:      userId,
		LotteryId:   req.LotteryId,
		PrizeName:   req.PrizeName,
		Content:     req.Content,
		Pics:        req.Pics,
		PraiseCount: 0,
	})
	if err != nil {
		return nil, err
	}
	return &types.CommentAddResp{}, nil
}
