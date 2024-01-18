package lottery

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLotteryLogic {
	return &UpdateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLotteryLogic) UpdateLottery(req *types.UpdateLotteryReq) (resp *types.UpdateLotteryResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.LotteryRpc.UpdateLottery(l.ctx, &lottery.UpdateLotteryReq{
		UserId:      userId,
		Id:          int64(req.Id),
		PublishTime: req.PublishTime,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("update lottery fail"), "update lottery rpc UpdateLottery fail req: %+v , err : %v ", req, err)
	}
	return &types.UpdateLotteryResp{}, nil
}
