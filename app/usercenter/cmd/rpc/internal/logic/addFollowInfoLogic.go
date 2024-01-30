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

type AddFollowInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFollowInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFollowInfoLogic {
	return &AddFollowInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户引流方式----------------------
func (l *AddFollowInfoLogic) AddFollowInfo(in *pb.AddUserFollowReq) (*pb.AddUserFollowResp, error) {
	FollowInfo := new(model.UserFollow)
	_ = copier.Copy(&FollowInfo, in)
	_, err := l.svcCtx.UserFollowModel.Insert(l.ctx, FollowInfo)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "usercenter Database Exception FollowInfo : %+v , err: %v", FollowInfo, err)
	}
	return &pb.AddUserFollowResp{}, nil
}
