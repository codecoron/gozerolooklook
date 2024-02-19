package userContact

import (
	"context"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMyContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyContactListLogic {
	return &MyContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyContactListLogic) MyContactList(req *types.MyContactListReq) (resp *types.MyContactListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
