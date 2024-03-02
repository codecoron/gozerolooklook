package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type LotteryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLotteryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LotteryDetailLogic {
	return &LotteryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LotteryDetailLogic) LotteryDetail(in *pb.LotteryDetailReq) (resp *pb.LotteryDetailResp, err error) {
	lotteryId := in.Id
	res, err := l.svcCtx.PrizeModel.FindByLotteryId(l.ctx, lotteryId)
	if err != nil {
		return nil, err
	}
	lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, lotteryId)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_FIND_LOTTERY_BYLOTTERYID_ERROR), "LotteryDetail, lotteryId:%v, error: %v", lotteryId, err)
	}
	resp = new(pb.LotteryDetailResp)
	resp.Lottery = new(pb.Lottery)
	_ = copier.Copy(resp.Lottery, lottery)
	resp.Lottery.AnnounceTime = lottery.AnnounceTime.Unix()
	resp.Lottery.PublishTime = lottery.PublishTime.Time.Unix()
	resp.Lottery.AwardDeadline = lottery.AwardDeadline.Unix()
	resp.Lottery.CreateTime = lottery.CreateTime.Unix()
	resp.Lottery.UpdateTime = lottery.UpdateTime.Unix()

	for _, p := range res {
		prize := new(pb.Prize)
		_ = copier.Copy(prize, p)
		resp.Prizes = append(resp.Prizes, prize)
	}

	// 获取当前用户是否参与当前lottery
	count, err := l.svcCtx.LotteryParticipationModel.CheckIsParticipatedByUserIdAndLotteryId(l.ctx, in.UserId, lotteryId)
	if err != nil {
		return nil, err
	}
	resp.IsParticipated = count
	return resp, nil
}
