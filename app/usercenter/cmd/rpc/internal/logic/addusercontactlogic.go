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

type AddUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserContactLogic {
	return &AddUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户联系方式----------------------
func (l *AddUserContactLogic) AddUserContact(in *pb.AddUserContactReq) (*pb.AddUserContactResp, error) {
	contact := new(model.UserContact)
	_ = copier.Copy(&contact, in)
	insert, err := l.svcCtx.UserContactModel.Insert(l.ctx, contact)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "usercenter Database Exception contact : %+v , err: %v", contact, err)
	}
	id, _ := insert.LastInsertId()
	return &pb.AddUserContactResp{Id: id}, nil
}
