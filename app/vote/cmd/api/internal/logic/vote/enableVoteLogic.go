package vote

import (
	"context"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnableVoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEnableVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnableVoteLogic {
	return &EnableVoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EnableVoteLogic) EnableVote(req *types.EnableVoteReq) (resp *types.EnableVoteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
