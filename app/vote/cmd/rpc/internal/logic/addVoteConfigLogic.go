package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/model"
	"looklook/common/xerr"

	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVoteConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVoteConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVoteConfigLogic {
	return &AddVoteConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------投票表-----------------------
func (l *AddVoteConfigLogic) AddVoteConfig(in *pb.AddVoteConfigReq) (*pb.AddVoteConfigResp, error) {
	voteConfig := new(model.VoteConfig)
	err := copier.Copy(voteConfig, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}

	builder := l.svcCtx.VoteConfigModel.SelectBuilder().Where("lottery_id = ?", voteConfig.LotteryId)
	voteData, err := l.svcCtx.VoteConfigModel.FindAll(l.ctx, builder, "")
	if voteData != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.VOTE_VOTE_RECORD_NOT_FOUND, fmt.Sprintf("LotteryId:%d 投票信息已存在,请勿重复创建!", voteConfig.LotteryId)), "voteConfig NOT FOUND err:%v, voteConfig:%+v", err)
	}

	insertResult, err := l.svcCtx.VoteConfigModel.Insert(l.ctx, voteConfig)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add voteConfig db vote_config Insert err:%v, address:%+v", err, voteConfig)
	}
	lastId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add address db user_address insertResult.LastInsertId err:%v, address:%+v", err, voteConfig)
	}

	return &pb.AddVoteConfigResp{
		Id: lastId,
	}, nil
}
