package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserFollowLogic {
	return &AddUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserFollowLogic) AddUserFollow(in *pb.AddUserFollowReq) (*pb.AddUserFollowResp, error) {
	Follow := new(model.UserFollow)
	_ = copier.Copy(&Follow, in)
	_, err := l.svcCtx.UserFollowModel.Insert(l.ctx, Follow)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "usercenter Database Exception Follow : %+v , err: %v", Follow, err)
	}
	return &pb.AddUserFollowResp{}, nil
}
