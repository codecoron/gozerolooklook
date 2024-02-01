package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserContactLogic {
	return &SearchUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserContactLogic) SearchUserContact(in *pb.SearchUserContactReq) (*pb.SearchUserContactResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserContactResp{}, nil
}
