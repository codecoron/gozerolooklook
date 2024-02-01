package userfollow

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *AddfollowInfoLogic) AddfollowInfo(req *types.AddfollowReq) (resp *types.AddfollowInfoResp, err error) {
	_, err = l.svcCtx.UsercenterRpc.AddUserFollow(l.ctx, &usercenter.AddUserFollowReq{
		UserId:     req.UserID,
		Type:       req.Type,
		AppletType: req.AppletType,
		Name:       req.Name,
		Desc:       req.Desc,
		Avatar:     req.Avatar,
		IsShow:     req.IsShow,
		QrCode:     req.QrCode,
		InputA:     req.InputA,
		InputB:     req.InputB,
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return
}
