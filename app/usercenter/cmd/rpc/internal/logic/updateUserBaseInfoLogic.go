package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserBaseInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserBaseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserBaseInfoLogic {
	return &UpdateUserBaseInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserBaseInfoLogic) UpdateUserBaseInfo(in *pb.UpdateUserBaseInfoReq) (*pb.UpdateUserBaseInfoResp, error) {
	// todo: 暂时没有不带session的update 使用事务写入
	err := l.svcCtx.UserModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "user not exist err:%v,user:%+v", err, user)
		}
		user.Nickname = in.Nickname
		user.Sex = in.Sex
		user.Info = in.Info
		user.Avatar = in.Avatar
		_, err = l.svcCtx.UserModel.Update(context, session, user)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "update user base info err:%v,user:%+v", err, user)
		}
		return nil
	})
	if err != nil {
		logx.Error("update user base info err:", err)
		return nil, err
	}
	return &pb.UpdateUserBaseInfoResp{}, nil
}
