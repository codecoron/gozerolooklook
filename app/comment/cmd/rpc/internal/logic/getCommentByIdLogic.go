package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByIdLogic {
	return &GetCommentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentByIdLogic) GetCommentById(in *pb.GetCommentByIdReq) (*pb.GetCommentByIdResp, error) {
	// todo: add your logic here and delete this line
	comment, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	pbComment := new(pb.Comment)

	err = copier.Copy(&pbComment, &comment)
	if err != nil {
		return nil, err
	}

	return &pb.GetCommentByIdResp{
		Comment: pbComment,
	}, nil
}
