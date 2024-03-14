package logic

import (
	"context"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCommentLogic {
	return &SearchCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCommentLogic) SearchComment(in *pb.SearchCommentReq) (*pb.SearchCommentResp, error) {
	if in.LastId == 0 {
		id, err := l.svcCtx.CommentModel.GetCommentLastId()
		if err != nil {
			return nil, err
		}
		in.LastId = id + 1
	}
	list, err := l.svcCtx.CommentModel.CommentList(l.ctx, in.Limit, in.LastId, in.Sort)
	if err != nil {
		return nil, err
	}
	var resp []*pb.Comment
	if len(list) > 0 {
		for _, comment := range list {
			var pbComment pb.Comment
			pbComment.Id = comment.Id
			pbComment.UserId = comment.UserId
			pbComment.LotteryId = comment.LotteryId
			pbComment.PrizeName = comment.PrizeName
			pbComment.Content = comment.Content
			pbComment.Pics = comment.Pics
			pbComment.PraiseCount = comment.PraiseCount
			pbComment.CreateTime = comment.CreateTime.Unix()
			pbComment.UpdateTime = comment.UpdateTime.Unix()
			resp = append(resp, &pbComment)
		}
	}

	return &pb.SearchCommentResp{
		Comment: resp,
	}, nil
}
