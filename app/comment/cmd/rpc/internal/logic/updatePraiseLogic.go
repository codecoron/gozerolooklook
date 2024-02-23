package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &pb.UpdatePraiseResp{}, nil
}
