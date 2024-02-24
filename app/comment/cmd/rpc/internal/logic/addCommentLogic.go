package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"
	"looklook/app/comment/model"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------comment-----------------------

// AddComment 添加评论
func (l *AddCommentLogic) AddComment(in *pb.AddCommentReq) (*pb.AddCommentResp, error) {
	comment := new(model.Comment)
	comment.UserId = in.UserId
	comment.LotteryId = in.LotteryId
	comment.PrizeName = in.PrizeName
	comment.Content = in.Content
	comment.Pics = in.Pics
	comment.PraiseCount = in.PraiseCount

	_, err := l.svcCtx.CommentModel.Insert(l.ctx, comment)
	if err != nil {
		return nil, err
	}

	return &pb.AddCommentResp{}, nil
}
