package logic

import (
	"context"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLastIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLastIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLastIdLogic {
	return &GetCommentLastIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLastIdLogic) GetCommentLastId(in *pb.GetCommentLastIdReq) (*pb.GetCommentLastIdResp, error) {
	id, err := l.svcCtx.CommentModel.GetCommentLastId()
	if err != nil {
		return nil, err
	}

	return &pb.GetCommentLastIdResp{
		LastId: id,
	}, nil
}
