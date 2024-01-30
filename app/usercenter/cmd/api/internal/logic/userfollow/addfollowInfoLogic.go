package userfollow

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type AddfollowInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddfollowInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddfollowInfoLogic {
	return &AddfollowInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddfollowInfoLogic) AddfollowInfo(req *types.AddfollowInfoReq) (resp *types.AddfollowInfoResp, err error) {
	var pbAddUserfollowReq usercenter.AddFollowInfoReq
	_ = copier.Copy(&pbAddUserfollowReq, req)
	info, err := l.svcCtx.UsercenterRpc.AddFollowInfo(l.ctx, &pbAddUserfollowReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	_ = copier.Copy(resp, &info)
	return &types.AddfollowInfoResp{}, nil
}
