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

type SearchUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserAddressLogic {
	return &SearchUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserAddressLogic) SearchUserAddress(in *pb.SearchUserAddressReq) (*pb.SearchUserAddressResp, error) {
	list, err := l.svcCtx.UserAddressModel.List(l.ctx, in.Page, in.Limit)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user's homestay order err : %v , in :%+v", err, in)
	}

	var resp []*pb.UserAddress
	if len(list) > 0 {
		for _, address := range list {
			var pbAddress pb.UserAddress
			_ = copier.Copy(&pbAddress, address)
			resp = append(resp, &pbAddress)
		}
	}

	return &pb.SearchUserAddressResp{
		UserAddress: resp,
	}, nil
}
