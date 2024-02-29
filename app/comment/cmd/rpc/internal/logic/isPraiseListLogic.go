package logic

import (
	"context"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsPraiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsPraiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsPraiseListLogic {
	return &IsPraiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsPraiseListLogic) IsPraiseList(in *pb.IsPraiseListReq) (*pb.IsPraiseListResp, error) {
	list, err := l.svcCtx.PraiseModel.IsPraiseList(l.ctx, in.CommentId, in.UserId)
	if err != nil {
		return nil, err
	}
	//ids := make([]int64, 0, len(list))
	//for _, v := range list {
	//	ids = append(ids, v)
	//}

	return &pb.IsPraiseListResp{
		PraiseId: list,
	}, nil
}
