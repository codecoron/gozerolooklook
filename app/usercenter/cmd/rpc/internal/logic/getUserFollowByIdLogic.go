package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFollowByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowByIdLogic {
	return &GetUserFollowByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFollowByIdLogic) GetUserFollowById(in *pb.GetUserFollowByIdReq) (*pb.GetUserFollowByIdResp, error) {

	one, err := l.svcCtx.UserFollowModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrMsg("get user follow  fail"), "err : %v , in : %+v", err, in)
	}

	var respUserFollow usercenter.UserFollow
	_ = copier.Copy(&respUserFollow, one)

	return &pb.GetUserFollowByIdResp{
		UserFollow: &respUserFollow,
	}, nil
}
