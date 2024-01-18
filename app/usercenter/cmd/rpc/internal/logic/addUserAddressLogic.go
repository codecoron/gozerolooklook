package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserAddressLogic {
	return &AddUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------用户收货地址表-----------------------
func (l *AddUserAddressLogic) AddUserAddress(in *pb.AddUserAddressReq) (*pb.AddUserAddressResp, error) {
	userAddress := new(model.UserAddress)
	err := copier.Copy(userAddress, in)
	if err != nil {
		//todo 优化错误码
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "copier : %+v , err: %v", in, err)
	}

	insertResult, err := l.svcCtx.UserAddressModel.Insert(l.ctx, userAddress)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add address db user_address Insert err:%v, address:%+v", err, userAddress)
	}
	lastId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Add address db user_address insertResult.LastInsertId err:%v, address:%+v", err, userAddress)
	}

	return &pb.AddUserAddressResp{
		Id: lastId,
	}, nil
}
