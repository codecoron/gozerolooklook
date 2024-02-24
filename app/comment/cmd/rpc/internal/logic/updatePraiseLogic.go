package logic

import (
	"context"
	"looklook/app/comment/model"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePraiseLogic {
	return &UpdatePraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePraiseLogic) UpdatePraise(in *pb.UpdatePraiseReq) (*pb.UpdatePraiseResp, error) {
	praise := new(model.Praise)
	praise.Id = in.Id
	praise.CommentId = in.CommentId
	praise.UserId = in.UserId

	err := l.svcCtx.PraiseModel.Update(l.ctx, praise)
	if err != nil {
		return nil, err
	}
	return &pb.UpdatePraiseResp{}, nil
}
