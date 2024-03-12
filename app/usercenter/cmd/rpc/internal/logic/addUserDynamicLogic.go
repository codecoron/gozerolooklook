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

type AddUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserDynamicLogic {
	return &AddUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserDynamicLogic) AddUserDynamic(in *pb.AddUserDynamicReq) (*pb.AddUserDynamicResp, error) {
	userDynamic := new(model.UserDynamic)
	err := copier.Copy(userDynamic, in)
	if err != nil {
		//todo 优化错误码
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}

	insertResult, err := l.svcCtx.UserDynamicModel.Insert(l.ctx, userDynamic)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add userDynamic db user_dynamic Insert err:%v, dynamic:%+v", err, userDynamic)
	}
	lastId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add userDynamic db user_dynamic insertResult.LastInsertId err:%v, dynamic:%+v", err, userDynamic)
	}

	return &pb.AddUserDynamicResp{
		Id: lastId,
	}, nil

	return &pb.AddUserDynamicResp{}, nil
}
