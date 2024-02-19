package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaskRecordByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskRecordByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskRecordByIdLogic {
	return &GetTaskRecordByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskRecordByIdLogic) GetTaskRecordById(in *pb.GetTaskRecordByIdReq) (*pb.GetTaskRecordByIdResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetTaskRecordByIdResp{}, nil
}
