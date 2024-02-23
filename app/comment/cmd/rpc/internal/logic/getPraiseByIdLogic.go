package logic

import (
	"context"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPraiseByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPraiseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPraiseByIdLogic {
	return &GetPraiseByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPraiseByIdLogic) GetPraiseById(in *pb.GetPraiseByIdReq) (*pb.GetPraiseByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetPraiseByIdResp{}, nil
}
