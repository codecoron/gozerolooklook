package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/comment/model"
	"looklook/common/xerr"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPraiseLogic {
	return &AddPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddPraise -----------------------praise-----------------------
func (l *AddPraiseLogic) AddPraise(in *pb.AddPraiseReq) (*pb.AddPraiseResp, error) {
	praise := new(model.Praise)
	praise.CommentId = in.CommentId
	praise.UserId = in.UserId

	_, err := l.svcCtx.PraiseModel.Insert(l.ctx, praise)
	if err != nil {
		return &pb.AddPraiseResp{}, errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTPRAISE_ERROR), "praise Database Exception praise : %+v , err: %v", praise, err)
	}

	return &pb.AddPraiseResp{}, nil
}
