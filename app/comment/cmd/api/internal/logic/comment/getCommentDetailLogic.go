package comment

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/comment/cmd/rpc/comment"

	"looklook/app/comment/cmd/api/internal/svc"
	"looklook/app/comment/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCommentDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentDetailLogic {
	return &GetCommentDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCommentDetailLogic) GetCommentDetail(req *types.CommentDetailReq) (resp *types.CommentDetailResp, err error) {
	res, err := l.svcCtx.CommentRpc.GetCommentById(l.ctx, &comment.GetCommentByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	commentDetail := new(types.Comment)
	err = copier.Copy(commentDetail, &res.Comment)
	if err != nil {
		return nil, err
	}

	return &types.CommentDetailResp{
		Comment: *commentDetail,
	}, nil
}
