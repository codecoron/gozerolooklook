package logic

import (
	"context"
	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsPraiseLogic {
	return &IsPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsPraiseLogic) IsPraise(in *pb.IsPraiseReq) (*pb.IsPraiseResp, error) {
	PraiseId, err := l.svcCtx.PraiseModel.IsPraise(l.ctx, in.CommentId, in.UserId)
	// todo 封装统一错误后怎么根据错误码进行处理
	if err != nil {
		return nil, err
	}
	return &pb.IsPraiseResp{PraiseId: PraiseId}, nil
}
