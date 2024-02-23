package logic

import (
	"context"
	"looklook/app/comment/model"
	"time"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
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
func (l *AddCommentLogic) AddComment(in *pb.AddCommentReq) (*pb.AddCommentResp, error) {
	_, err := l.svcCtx.CommentModel.Insert(l.ctx, &model.Comment{
		UserId:      in.UserId,
		LotteryId:   in.LotteryId,
		PrizeName:   in.PrizeName,
		Content:     in.Content,
		Pics:        in.Pics,
		PraiseCount: in.PraiseCount,
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddCommentResp{}, nil
}
