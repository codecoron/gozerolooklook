package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLotteryParticipationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLotteryParticipationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLotteryParticipationLogic {
	return &SearchLotteryParticipationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLotteryParticipationLogic) SearchLotteryParticipation(in *pb.SearchLotteryParticipationReq) (*pb.SearchLotteryParticipationResp, error) {
	offset := (in.PageIndex - 1) * in.PageSize
	limit := in.PageSize
	// todo 有sql语法错误
	builder := l.svcCtx.LotteryParticipationModel.SelectBuilder().Where("lottery_id = ?", in.LotteryId).Limit(uint64(limit)).Offset(uint64(offset))

	list, err := l.svcCtx.LotteryParticipationModel.FindAll(l.ctx, builder, "")
	if err != nil {
		return nil, err
	}

	resp := &pb.SearchLotteryParticipationResp{
		Count: int64(len(list)),
		List:  []*pb.LotteryParticipation{},
	}

	if err = copier.Copy(&resp.List, list); err != nil {
		return nil, err
	}

	return resp, nil
}
