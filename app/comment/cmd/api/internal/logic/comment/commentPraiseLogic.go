package comment

import (
	"context"
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/common/ctxdata"

	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentPraiseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentPraiseLogic {
	return &CommentPraiseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentPraiseLogic) CommentPraise(req *types.CommentPraiseReq) (resp *types.CommentPraiseResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.CommentRpc.PraiseComment(l.ctx, &comment.PraiseCommentReq{
		UserId:    userId,
		CommentId: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.CommentPraiseResp{}, err
}
