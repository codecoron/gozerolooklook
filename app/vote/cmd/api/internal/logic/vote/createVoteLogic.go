package vote

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"
	"looklook/app/vote/cmd/rpc/pb"
	"looklook/common/ctxdata"

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
	AddVoteConfigReq := new(pb.AddVoteConfigReq)
	err = copier.Copy(AddVoteConfigReq, req)
	if err != nil {
		return nil, err
	}
	AddVoteConfigReq.UserId = ctxdata.GetUidFromCtx(l.ctx)
	VoteConfigByte, err := json.Marshal(req.VoteConfig)
	if err != nil {
		return nil, err
	}
	AddVoteConfigReq.VoteConfig = string(VoteConfigByte)

	addVoteConfig, err := l.svcCtx.VoteRpc.AddVoteConfig(l.ctx, AddVoteConfigReq)
	if err != nil {
		return nil, errors.Wrapf(err, "add vote_config rpc fail req: %+v , err : %v ", req, err)
	}

	return &types.CreateVoteResp{Id: addVoteConfig.Id}, nil
}
