package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPrizeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPrizeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPrizeByIdLogic {
	return &GetPrizeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPrizeByIdLogic) GetPrizeById(in *pb.GetPrizeByIdReq) (*pb.GetPrizeByIdResp, error) {
	one, err := l.svcCtx.PrizeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	prize := &pb.Prize{}
	err = copier.Copy(prize, one)
	if err != nil {
		return nil, err
	}

	return &pb.GetPrizeByIdResp{
		Prize: prize,
	}, nil
}
