package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserFollowLogic {
	return &SearchUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserFollowLogic) SearchUserFollow(in *pb.SearchUserFollowReq) (*pb.SearchUserFollowResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SearchUserFollowResp{}, nil
}
