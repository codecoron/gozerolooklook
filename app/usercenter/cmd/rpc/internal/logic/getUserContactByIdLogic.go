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

var ErrContactNoExistsError = xerr.NewErrMsg("联系方式不存在")

type GetUserContactByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserContactByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserContactByIdLogic {
	return &GetUserContactByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserContactByIdLogic) GetUserContactById(in *pb.GetUserContactByIdReq) (*pb.GetUserContactByIdResp, error) {
	user, err := l.svcCtx.UserContactModel.FindOne(l.ctx, in.Id)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetUserContactById find contact db err , id:%d , err:%v", in.Id, err)
	}
	if user == nil {
		return nil, errors.Wrapf(ErrContactNoExistsError, "id:%d", in.Id)
	}
	var respUserContact usercenter.UserContact
	_ = copier.Copy(&respUserContact, user)

	return &pb.GetUserContactByIdResp{
		UserContact: &respUserContact,
	}, nil
}
