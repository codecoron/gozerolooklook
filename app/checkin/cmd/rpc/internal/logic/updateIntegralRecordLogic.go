package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIntegralRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIntegralRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIntegralRecordLogic {
	return &UpdateIntegralRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateIntegralRecordLogic) UpdateIntegralRecord(in *pb.UpdateIntegralRecordReq) (*pb.UpdateIntegralRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UpdateIntegralRecordResp{}, nil
}
