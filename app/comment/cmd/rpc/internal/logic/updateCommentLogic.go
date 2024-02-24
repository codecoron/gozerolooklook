package logic

import (
	"context"
	"looklook/app/comment/model"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) UpdateComment(in *pb.UpdateCommentReq) (*pb.UpdateCommentResp, error) {
	// todo: add your logic here and delete this line
	comment := new(model.Comment)
	comment.Id = in.Id
	comment.UserId = in.UserId
	comment.Content = in.Content
	comment.PrizeName = in.PrizeName
	comment.UserId = in.UserId
	comment.PraiseCount = in.PraiseCount
	comment.Pics = in.Pics

	err := l.svcCtx.CommentModel.Update(l.ctx, comment)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateCommentResp{}, nil
}
