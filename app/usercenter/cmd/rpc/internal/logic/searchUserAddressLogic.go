package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserAddressLogic {
	return &SearchUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserAddressLogic) SearchUserAddress(in *pb.SearchUserAddressReq) (*pb.SearchUserAddressResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserAddressResp{}, nil
}
