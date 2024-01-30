package userfollow

import (
	"context"
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
	//if req == nil {
	//	return nil, errors.Wrapf(err, "req: %+v", req)
	//}
	var pbReq *usercenter.AddUserFollowReq
	//_ = copier.Copy(&pbReq, req)
	//var pbReq *usercenter.AddUserFollowReq
	pbReq.Name = req.Name
	pbReq.Desc = req.Desc
	pbReq.QrCode = req.QrCode
	pbReq.UserId = 1
	pbReq.Avatar = req.Avatar
	pbReq.QrCode = req.QrCode
	pbReq.InputA = req.InputA

	_, err = l.svcCtx.UsercenterRpc.AddUserFollow(l.ctx, pbReq)
	if err != nil {
		return nil, errors.Wrapf(err, "req: %+v", req)
	}
	return &types.AddfollowInfoResp{}, nil
}
