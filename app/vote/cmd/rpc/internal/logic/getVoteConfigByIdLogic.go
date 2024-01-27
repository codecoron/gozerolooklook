package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVoteConfigByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVoteConfigByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVoteConfigByIdLogic {
	return &GetVoteConfigByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetVoteConfigByIdLogic) GetVoteConfigById(in *pb.GetVoteConfigByIdReq) (resp *pb.GetVoteConfigByIdResp, err error) {
	voteConfigId := in.Id
	voteConfig, err := l.svcCtx.VoteConfigModel.FindOne(l.ctx, voteConfigId)
	if err != nil {
		return nil, err
	}

	resp = new(pb.GetVoteConfigByIdResp)
	resp.VoteConfig = new(pb.VoteConfig)

	// 使用 copier.Copy 进行拷贝
	if err := copier.Copy(resp.VoteConfig, voteConfig); err != nil {
		return nil, err
	}

	return resp, nil
}
