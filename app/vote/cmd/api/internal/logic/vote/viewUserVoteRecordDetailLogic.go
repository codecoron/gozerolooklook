package vote

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/vote/cmd/rpc/vote"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/vote/cmd/api/internal/svc"
	"looklook/app/vote/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewUserVoteRecordDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewViewUserVoteRecordDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewUserVoteRecordDetailLogic {
	return &ViewUserVoteRecordDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ViewUserVoteRecordDetailLogic) ViewUserVoteRecordDetail(req *types.ViewUserVoteRecordDetailReq) (resp *types.ViewUserVoteRecordDetailResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.VoteRpc.GetUserVoteRecordDetail(l.ctx, &vote.GetUserVoteRecordDetailReq{
		LotteryId: req.LotteryId,
		UserId:    userId,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "viewUserVoteRecordDetail rpc fail req: %+v , err : %v ", req, err)
	}

	//fmt.Println("----res----", res)

	resp = &types.ViewUserVoteRecordDetailResp{}

	if err := copier.CopyWithOption(resp, res.UserVoteRecordDetails, copier.Option{
		DeepCopy: true,
	}); err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to copy VoteRecord to ViewVoteDetailResp"), "Failed to copy VoteRecord to ViewVoteRecordDetailResp err: %v", err)
	}

	for _, userVoteRecordDetail := range res.UserVoteRecordDetails {
		resp.VoteUserRecordDetails = append(resp.VoteUserRecordDetails, &types.VoteUserRecordDetail{
			LotteryId:      userVoteRecordDetail.LotteryId,
			UserId:         userVoteRecordDetail.UserId,
			SelectedOption: userVoteRecordDetail.SelectedOption,
		})
	}

	//fmt.Println("----resp----", resp)

	return resp, nil
}
