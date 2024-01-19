package address

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConvertAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConvertAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConvertAddressLogic {
	return &ConvertAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConvertAddressLogic) ConvertAddress(req *types.ConvertAddressReq) (resp *types.ConvertAddressResp, err error) {
	// todo: add your logic here and delete this line

	return
}
