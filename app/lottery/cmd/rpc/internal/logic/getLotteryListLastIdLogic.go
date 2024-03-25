package logic

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLotteryListLastIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLotteryListLastIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLotteryListLastIdLogic {
	return &GetLotteryListLastIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLotteryListLastIdLogic) GetLotteryListLastId(in *pb.GetLotteryListLastIdReq) (*pb.GetLotteryListLastIdResp, error) {
	id, err := l.svcCtx.LotteryModel.GetLastId(l.ctx)
	if err != nil {
		return nil, err
	}
	return &pb.GetLotteryListLastIdResp{LastId: id}, nil
}
