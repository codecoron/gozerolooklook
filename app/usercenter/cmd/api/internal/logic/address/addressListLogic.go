package address

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddressListLogic {
	return &AddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddressListLogic) AddressList(req *types.AddressListReq) (resp *types.AddressListResp, err error) {
	rpcAddressList, err := l.svcCtx.UsercenterRpc.SearchUserAddress(l.ctx, &usercenter.SearchUserAddressReq{
		Page:  req.Page,
		Limit: req.PageSize,
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get SearchLottery"), "Failed to get SearchLottery err : %v ,req:%+v", err, req)
	}

	var AddressList []types.UserAddress
	if len(rpcAddressList.UserAddress) > 0 {
		for _, item := range rpcAddressList.UserAddress {
			var t types.UserAddress
			_ = copier.Copy(&t, item)
			_ = json.Unmarshal([]byte(item.District), &t.District)
			AddressList = append(AddressList, t)
		}
	}

	return &types.AddressListResp{List: AddressList}, nil
}
