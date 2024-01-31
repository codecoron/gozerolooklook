package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsParticipatedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsParticipatedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsParticipatedLogic {
	return &CheckIsParticipatedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckIsParticipatedLogic) CheckIsParticipated(in *pb.CheckIsParticipatedReq) (*pb.CheckIsParticipatedResp, error) {
	// 获取当前用户是否已参与当前抽奖
	resp := new(pb.CheckIsParticipatedResp)
	_, err := l.svcCtx.LotteryParticipationModel.FindOneByLotteryIdUserId(l.ctx, in.LotteryId, in.UserId)
	if err == sqlx.ErrNotFound {
		resp.IsParticipated = 0
	} else if err != nil {
		return nil, err
	} else {
		resp.IsParticipated = 1
	}

	return resp, nil
}
