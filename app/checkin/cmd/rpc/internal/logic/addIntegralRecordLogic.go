package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddIntegralRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddIntegralRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddIntegralRecordLogic {
	return &AddIntegralRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------integralRecord-----------------------
func (l *AddIntegralRecordLogic) AddIntegralRecord(in *pb.AddIntegralRecordReq) (*pb.AddIntegralRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddIntegralRecordResp{}, nil
}
