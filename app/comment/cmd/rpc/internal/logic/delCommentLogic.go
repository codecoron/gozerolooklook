package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelCommentLogic {
	return &DelCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelCommentLogic) DelComment(in *pb.DelCommentReq) (*pb.DelCommentResp, error) {
	// todo : 软删除评论
	// 删除评论
	//err := l.svcCtx.CommentModel.Delete(l.ctx, in.Id)
	data, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FINDCOMMENT_ERROR), "comment Database Exception commentId : %d , err: %v", in.Id, err)
	}
	err = l.svcCtx.CommentModel.DeleteSoft(l.ctx, data)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_DELETECOMMENT_ERROR), "comment Database Exception commentId : %d , err: %v", in.Id, err)
	}

	return &pb.DelCommentResp{}, nil
}
