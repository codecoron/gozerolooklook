package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIntegralRecordByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIntegralRecordByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntegralRecordByUserIdLogic {
	return &GetIntegralRecordByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIntegralRecordByUserIdLogic) GetIntegralRecordByUserId(in *pb.GetIntegralRecordByUserIdReq) (*pb.GetIntegralRecordByUserIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetIntegralRecordByUserIdResp{}, nil
}
