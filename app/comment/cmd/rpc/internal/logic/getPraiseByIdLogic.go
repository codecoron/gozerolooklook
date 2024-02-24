package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPraiseByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPraiseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPraiseByIdLogic {
	return &GetPraiseByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPraiseByIdLogic) GetPraiseById(in *pb.GetPraiseByIdReq) (*pb.GetPraiseByIdResp, error) {
	// todo: add your logic here and delete this line
	praise, err := l.svcCtx.PraiseModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	pbPraise := new(pb.Praise)

	err = copier.Copy(&pbPraise, &praise)
	if err != nil {
		return nil, err
	}

	return &pb.GetPraiseByIdResp{
		Praise: pbPraise,
	}, nil
}
