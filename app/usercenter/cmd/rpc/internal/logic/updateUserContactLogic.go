package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserContactLogic {
	return &UpdateUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserContactLogic) UpdateUserContact(in *pb.UpdateUserContactReq) (*pb.UpdateUserContactResp, error) {
	// todo: add your logic here and delete this line

	contact, err := l.svcCtx.UserContactModel.FindOne(l.ctx, in.Id)
	if err != nil {
		logx.Error("查询用户联系方式id失败:%v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "err: %v", err)
	}

	contact.Id = in.Id
	contact.Content = in.Content
	contact.Remark = in.Remark
	contactId, err := l.svcCtx.UserContactModel.UpDateUserContactById(l.ctx, contact.Id, contact.Content, contact.Remark)
	if err != nil {
		logx.Error("用户联系方式修改失败:%v", err)
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Contact Database Exception contact : %+v , err: %v", contact, err)
	}
	return &pb.UpdateUserContactResp{Id: contactId}, nil
}
