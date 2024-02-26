package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/comment/model"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PraiseCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPraiseCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PraiseCommentLogic {
	return &PraiseCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PraiseCommentLogic) PraiseComment(in *pb.PraiseCommentReq) (*pb.PraiseCommentResp, error) {
	// 评论点赞/取消点赞
	// 先判断是否已经点赞
	praiseId, err := l.svcCtx.PraiseModel.IsPraise(l.ctx, in.CommentId, in.UserId)
	if err != nil {
		return nil, err
	}
	// 如果已经点赞，则取消点赞
	if praiseId != 0 {
		err := l.svcCtx.CommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			err := l.svcCtx.PraiseModel.Delete(l.ctx, praiseId)
			if err != nil {
				return err
			}
			_, err = l.svcCtx.CommentModel.UpdatePraiseNum(l.ctx, in.CommentId, -1)
			if err != nil {
				return err
			}
			return nil

		})
		if err != nil {
			return nil, err
		}
		return &pb.PraiseCommentResp{}, nil
	} else {
		// 如果没有点赞，则点赞
		err := l.svcCtx.CommentModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
			praise := &model.Praise{
				CommentId: in.CommentId,
				UserId:    in.UserId,
			}
			_, err := l.svcCtx.PraiseModel.Insert(l.ctx, praise)
			if err != nil {
				return err
			}
			_, err = l.svcCtx.CommentModel.UpdatePraiseNum(l.ctx, in.CommentId, 1)
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return &pb.PraiseCommentResp{}, nil
}
