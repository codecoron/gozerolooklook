package logic

import (
	"context"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCommentLogic {
	return &GetUserCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserComment 获取当前用户评论列表
func (l *GetUserCommentLogic) GetUserComment(in *pb.GetUserCommentReq) (*pb.GetUserCommentResp, error) {
	// 获取当前用户评论列表，以及评论的点赞数，当前用户是否点赞
	// 1. 获取当前用户评论列表
	logx.Error(in)
	builder := l.svcCtx.CommentModel.SelectBuilder().Where("user_id = ?", in.UserId)
	logx.Error(&builder)
	list, err := l.svcCtx.CommentModel.FindAll(l.ctx, builder, "")
	logx.Error(&list)
	// 得到评论的id列表
	var commentIds []int64
	for _, v := range list {
		commentIds = append(commentIds, v.Id)
	}

	// 3. 获取当前用户是否点赞
	likeList, err := l.svcCtx.PraiseModel.IsPraiseList(l.ctx, commentIds, in.UserId)
	if err != nil {
		return nil, err
	}

	// 转成map
	likeMap := make(map[int64]int64)
	// 默认没有点赞
	for _, v := range commentIds {
		likeMap[v] = 0
	}

	for _, v := range likeList {
		likeMap[v] = 1
	}

	// 4. 组装返回数据
	var Comment []*pb.Comment
	for _, v := range list {
		Comment = append(Comment, &pb.Comment{
			Id:          v.Id,
			UserId:      v.UserId,
			Content:     v.Content,
			LotteryId:   v.LotteryId,
			PrizeName:   v.PrizeName,
			Pics:        v.Pics,
			PraiseCount: v.PraiseCount,
			IsPraise:    likeMap[v.Id],
		})
	}

	// 5. 返回数据
	return &pb.GetUserCommentResp{
		Comment: Comment,
	}, nil
}
