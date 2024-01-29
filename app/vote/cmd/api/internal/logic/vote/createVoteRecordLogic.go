package vote

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/cmd/rpc/pb"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateVoteRecordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateVoteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateVoteRecordLogic {
	return &CreateVoteRecordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateVoteRecordLogic) CreateVoteRecord(req *types.CreateVoteRecordReq) (resp *types.CreateVoteRecordResp, err error) {
	AddVoteRecordReq := new(pb.AddVoteRecordReq)
	err = copier.Copy(AddVoteRecordReq, req)
	if err != nil {
		return nil, err
	}
	AddVoteRecordReq.UserId = ctxdata.GetUidFromCtx(l.ctx)

	addVoteRecord, err := l.svcCtx.VoteRpc.AddVoteRecord(l.ctx, AddVoteRecordReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("add vote_record fail"), "add vote_record rpc AddVoteRecord fail req: %+v , err : %v ", req, err)
	}

	return &types.CreateVoteRecordResp{Id: addVoteRecord.Id}, nil
}
