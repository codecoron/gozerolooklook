package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchIntegralRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchIntegralRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchIntegralRecordLogic {
	return &SearchIntegralRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchIntegralRecordLogic) SearchIntegralRecord(in *pb.SearchIntegralRecordReq) (*pb.SearchIntegralRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchIntegralRecordResp{}, nil
}
