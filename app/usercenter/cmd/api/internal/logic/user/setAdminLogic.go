package user

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
	"looklook/app/usercenter/cmd/rpc/usercenter"
)

type SetAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetAdminLogic {
	return &SetAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetAdminLogic) SetAdmin(req *types.SetAdminReq) (resp *types.SetAdminResp, err error) {
	// 1 确定是否有管理员权限
	//ctxdata.GetUidFromCtx(l.ctx)
	//res, err := l.svcCtx.UsercenterRpc.CheckIsAdmin(l.ctx, &usercenter.CheckIsAdminReq{
	//	UserId: ctxdata.GetUidFromCtx(l.ctx),
	//})
	//if err != nil {
	//	return nil, err
	//}
	//if !res.IsAdmin {
	//	return nil, errors.New("没有操作权限")
	//}
	// 2 有则才能设置管理员
	_, err = l.svcCtx.UsercenterRpc.SetAdmin(l.ctx, &usercenter.SetAdminReq{
		UserId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return
}
