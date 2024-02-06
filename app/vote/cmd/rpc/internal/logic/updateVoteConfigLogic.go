package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"
	"looklook/app/vote/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVoteConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVoteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVoteConfigLogic {
	return &UpdateVoteConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVoteConfigLogic) UpdateVoteConfig(in *pb.UpdateVoteConfigReq) (*pb.UpdateVoteConfigResp, error) {
	voteConfig, err := l.svcCtx.VoteConfigModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "voteConfig err:%v,voteConfig:%+v", err, voteConfig)
	}
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.VOTE_VOTE_CONFIG_NOT_FOUND, "投票信息不存在!"), "voteConfig NOT FOUND err:%v,voteConfig:%+v", err, voteConfig)
	}
	err = copier.Copy(voteConfig, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}

	err = l.svcCtx.VoteConfigModel.Update(l.ctx, voteConfig)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "failed to update voteConfig err:%v,voteConfig:%+v", err, voteConfig)
	}

	return &pb.UpdateVoteConfigResp{}, nil
}
