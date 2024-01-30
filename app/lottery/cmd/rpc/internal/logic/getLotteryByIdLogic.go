package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/model"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryByIdLogic {
	return &GetLotteryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryByIdLogic) GetLotteryById(in *pb.GetLotteryByIdReq) (*pb.GetLotteryByIdResp, error) {
	lotteryInfo, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "lottery_id:%d,err:%v", in.Id, err)
	}

	pbLottery := new(pb.Lottery)
	_ = copier.Copy(pbLottery, lotteryInfo)

	return &pb.GetLotteryByIdResp{Lottery: pbLottery}, nil
}
