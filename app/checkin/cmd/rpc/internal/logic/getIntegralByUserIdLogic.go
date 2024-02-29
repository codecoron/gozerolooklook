package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIntegralByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIntegralByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIntegralByUserIdLogic {
	return &GetIntegralByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetIntegralByUserIdLogic) GetIntegralByUserId(in *pb.GetIntegralByUserIdReq) (*pb.GetIntegralByUserIdResp, error) {
	one, err := l.svcCtx.IntegralModel.FindOneByUserId(l.ctx, in.UserId)
	// 如果没有找到，说明用户还没使用过签到服务，返回0
	if err == sqlc.ErrNotFound {
		return &pb.GetIntegralByUserIdResp{
			Integral: 0,
		}, nil
	} else if err != nil {
		return nil, err
	}
	return &pb.GetIntegralByUserIdResp{
		Integral: one.Integral,
	}, nil
}
