package vote

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/cmd/rpc/vote"
	"looklook/common/xerr"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewVoteDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewVoteDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewVoteDetailLogic {
	return &ViewVoteDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewVoteDetailLogic) ViewVoteDetail(req *types.ViewVoteDetailReq) (resp *types.ViewVoteDetailResp, err error) {
	res, err := l.svcCtx.VoteRpc.GetVoteConfigById(l.ctx, &vote.GetVoteConfigByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "viewVoteDetail rpc fail req: %+v , err : %v ", req, err)
	}

	// 解析 VoteConfig 字段
	var voteConfig types.VoteConfigJSONData
	if err := json.Unmarshal([]byte(res.VoteConfig.VoteConfig), &voteConfig); err != nil {
		return nil, err
	}

	resp = &types.ViewVoteDetailResp{}
	if err := copier.Copy(resp, res.VoteConfig); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to copy VoteConfig to ViewVoteDetailResp"), "Failed to copy VoteConfig to ViewVoteDetailResp err : %v", err)
	}
	resp.VoteConfig = voteConfig

	return resp, nil
}
