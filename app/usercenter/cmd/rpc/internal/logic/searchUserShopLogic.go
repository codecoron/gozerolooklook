package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserShopLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserShopLogic {
	return &SearchUserShopLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserShopLogic) SearchUserShop(in *pb.SearchUserShopReq) (*pb.SearchUserShopResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserShopResp{}, nil
}
