package vote

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateVoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVoteLogic {
	return &UpdateVoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateVoteLogic) UpdateVote(req *types.UpdateVoteReq) (resp *types.UpdateVoteResp, err error) {
	UpdateVoteConfigReq := new(pb.UpdateVoteConfigReq)
	err = copier.Copy(UpdateVoteConfigReq, req)
	if err != nil {
		return nil, err
	}

	//解析转移json字符串
	VoteConfigByte, err := json.Marshal(req.VoteConfig)
	if err != nil {
		return nil, err
	}
	UpdateVoteConfigReq.VoteConfig = string(VoteConfigByte)

	_, err = l.svcCtx.VoteRpc.UpdateVoteConfig(l.ctx, UpdateVoteConfigReq)
	if err != nil {
		return nil, errors.Wrapf(err, "add vote_config rpc fail req: %+v , err : %v ", req, err)
	}
	return
}
