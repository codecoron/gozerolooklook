package comment

import (
	"context"
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

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
	// 如果用户没有中奖，不允许评论
	isWon, err := l.svcCtx.LotteryRpc.CheckUserIsWon(l.ctx, &lottery.CheckUserIsWonReq{
		LotteryId: req.LotteryId,
		UserId:    userId,
	})
	if err != nil {
		return nil, err
	}
	// 判断是否中奖
	if isWon.IsWon == 0 {
		return nil, xerr.NewErrCode(xerr.ErrUserNotWon)
	}

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
