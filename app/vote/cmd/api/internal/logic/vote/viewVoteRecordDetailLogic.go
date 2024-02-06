package vote

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/cmd/rpc/vote"
	"looklook/common/xerr"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewVoteRecordDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewVoteRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewVoteRecordDetailLogic {
	return &ViewVoteRecordDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewVoteRecordDetailLogic) ViewVoteRecordDetail(req *types.ViewVoteRecordDetailReq) (resp *types.ViewVoteRecordDetailResp, err error) {
	res, err := l.svcCtx.VoteRpc.GetVoteRecordDetail(l.ctx, &vote.GetVoteRecordDetailReq{
		LotteryId: req.LotteryId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "VoteRecordDetail rpc fail req: %+v , err : %v ", req, err)
	}

	//fmt.Println("----res----", res)

	resp = &types.ViewVoteRecordDetailResp{}
	if err := copier.CopyWithOption(resp, res, copier.Option{
		DeepCopy: true,
	}); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to copy VoteRecord to ViewVoteDetailResp"), "Failed to copy VoteRecord to ViewVoteRecordDetailResp err: %v", err)
	}

	//fmt.Println("----resp----", resp)

	return resp, nil
}
