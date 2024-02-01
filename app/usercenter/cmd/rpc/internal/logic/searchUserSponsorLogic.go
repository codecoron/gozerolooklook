package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserSponsorLogic {
	return &SearchUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserSponsorLogic) SearchUserSponsor(in *pb.SearchUserSponsorReq) (*pb.SearchUserSponsorResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserSponsorResp{}, nil
}
