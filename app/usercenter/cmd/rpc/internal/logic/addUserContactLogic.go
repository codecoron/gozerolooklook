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

// -----------------------抽奖发起人联系方式-----------------------
func (l *AddUserContactLogic) AddUserContact(in *pb.AddUserContactReq) (*pb.AddUserContactResp, error) {
	userContact := new(model.UserContact)
	err := copier.Copy(userContact, in)
	if err != nil {
		//todo 优化错误码
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}
	insert, err := l.svcCtx.UserContactModel.Insert(l.ctx, userContact)
	if err != nil {
		return nil, err
	}
	lastId, err := insert.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add Contact db user_Contact insertResult.LastInsertId err:%v, Contact:%+v", err, userContact)
	}
	return &pb.AddUserContactResp{
		Id: lastId,
	}, nil
}
