package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchTaskRecordLogic {
	return &SearchTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchTaskRecordLogic) SearchTaskRecord(in *pb.SearchTaskRecordReq) (*pb.SearchTaskRecordResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchTaskRecordResp{}, nil
}
