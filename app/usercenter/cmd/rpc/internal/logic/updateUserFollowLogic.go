package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserFollowLogic {
	return &UpdateUserFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserFollowLogic) UpdateUserFollow(in *pb.UpdateUserFollowReq) (*pb.UpdateUserFollowResp, error) {
	err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		one, err := l.svcCtx.UserFollowModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "userfollow not exist err:%v,userfollow:%+v", err, one)
		}
		one.Name = in.Name
		one.Type = in.Type
		one.Avatar = in.Avatar
		one.AppletType = in.AppletType
		one.Desc = in.Desc
		one.InputA = in.InputA
		one.InputB = in.InputB
		one.IsShow = in.IsShow
		one.QrCode = in.QrCode
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update userfollow base info err:%v,userfollow:%+v", err, one)
		}
		return nil
	})
	if err != nil {
		logx.Error("update userfollow base info err:", err)
		return nil, err
	}

	return &pb.UpdateUserFollowResp{}, nil
}
