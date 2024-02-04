package logic

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeWishCheckinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeWishCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeWishCheckinLogic {
	return &NoticeWishCheckinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------others----------------------
func (l *NoticeWishCheckinLogic) NoticeWishCheckin(in *pb.NoticeWishCheckinReq) (*pb.NoticeWishCheckinResp, error) {
	// todo: add your logic here and delete this line

	return &pb.NoticeWishCheckinResp{}, nil
}
