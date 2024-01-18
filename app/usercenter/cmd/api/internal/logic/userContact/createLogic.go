package userContact

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateReq) (resp *types.CreateResp, err error) {
	var pbAddUserContactReq usercenter.AddUserContactReq
	_ = copier.Copy(&pbAddUserContactReq, req)
	pbAddUserContactReq.UserId = ctxdata.GetUidFromCtx(l.ctx)
	pbRes, err := l.svcCtx.UsercenterRpc.AddUserContact(l.ctx, &pbAddUserContactReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.CreateResp{
		Id: pbRes.Id,
	}, nil
}
