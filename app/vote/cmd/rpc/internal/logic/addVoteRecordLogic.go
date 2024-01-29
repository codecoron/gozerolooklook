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
