package user

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateLogic) Update(req *types.UserUpdateReq) (resp *types.UserUpdateResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	_, err = l.svcCtx.UsercenterRpc.UpdateUserBaseInfo(l.ctx, &usercenter.UpdateUserBaseInfoReq{
		Id:        userId,
		Nickname:  req.Nickname,
		Sex:       req.Sex,
		Info:      req.Info,
		Avatar:    req.Avatar,
		Signature: req.Signature,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return
}
