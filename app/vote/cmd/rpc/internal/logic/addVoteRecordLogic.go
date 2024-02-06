package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/cmd/rpc/internal/svc"
	"looklook/app/vote/cmd/rpc/pb"
	"looklook/app/vote/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVoteRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddVoteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVoteRecordLogic {
	return &AddVoteRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------投票记录表-----------------------
func (l *AddVoteRecordLogic) AddVoteRecord(in *pb.AddVoteRecordReq) (*pb.AddVoteRecordResp, error) {
	voteRecord := new(model.VoteRecord)
	err := copier.Copy(voteRecord, in)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}

	builder := l.svcCtx.VoteConfigModel.SelectBuilder().Where("lottery_id = ?", voteRecord.LotteryId)
	results, err := l.svcCtx.VoteConfigModel.FindAll(l.ctx, builder, "")

	//fmt.Println("rows---------", row)

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.DB_ERROR, "FindAll voteConfig Fail"), "FindAll voteConfig db vote_config err:%v, voteConfig:%+v", err, voteRecord)
	}

	if len(results) == 0 {
		return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.VOTE_VOTE_RECORD_NOT_FOUND, fmt.Sprintf("LotteryId:%d 投票信息不存在,请添加配置后再投票!", voteRecord.LotteryId)), "voteConfig NOT FOUND, voteConfig:%+v", voteRecord)
	}

	//检查选项与selectedOption是否一致
	//pb.selectedOption
	//pb.VoteConfig
	getVoteConfig := new(pb.GetVoteConfigByIdResp)
	getVoteConfig.VoteConfig = new(pb.VoteConfig)

	//fmt.Println("results:-------", results)
	//fmt.Println("config:-------", getVoteConfig.VoteConfig)

	insertResult, err := l.svcCtx.VoteRecordModel.Insert(l.ctx, voteRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add voteRecord db vote_record Insert err:%v, voteRecord:%+v", err, voteRecord)
	}

	lastId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add voteRecord db vote_record insertResult.LastInsertId err:%v, voteRecord:%+v", err, voteRecord)
	}

	return &pb.AddVoteRecordResp{
		Id: lastId,
	}, nil

}
