package vote

import (
	"context"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVoteLogic {
	return &CreateVoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVoteLogic) CreateVote(req *types.CreateVoteReq) (resp *types.CreateVoteResp, err error) {
	// todo: add your logic here and delete this line

	return
}
