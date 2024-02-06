package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIntegralRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIntegralRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntegralRecordByIdLogic {
	return &GetIntegralRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIntegralRecordByIdLogic) GetIntegralRecordById(in *pb.GetIntegralRecordByIdReq) (*pb.GetIntegralRecordByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetIntegralRecordByIdResp{}, nil
}
