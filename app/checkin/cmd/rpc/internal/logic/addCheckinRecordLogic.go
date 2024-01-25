package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCheckinRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCheckinRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCheckinRecordLogic {
	return &AddCheckinRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------checkinRecord-----------------------
func (l *AddCheckinRecordLogic) AddCheckinRecord(in *pb.AddCheckinRecordReq) (*pb.AddCheckinRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.AddCheckinRecordResp{}, nil
}
