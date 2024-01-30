package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelIntegralRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelIntegralRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelIntegralRecordLogic {
	return &DelIntegralRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelIntegralRecordLogic) DelIntegralRecord(in *pb.DelIntegralRecordReq) (*pb.DelIntegralRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.DelIntegralRecordResp{}, nil
}
