package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelTaskRecordLogic {
	return &DelTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelTaskRecordLogic) DelTaskRecord(in *pb.DelTaskRecordReq) (*pb.DelTaskRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelTaskRecordResp{}, nil
}
