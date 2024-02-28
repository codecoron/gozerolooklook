package userContact

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpDateContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpDateContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpDateContactLogic {
	return &UpDateContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpDateContactLogic) UpDateContact(req *types.UpDateContactReq) (resp *types.UpDateContactResp, err error) {
	// todo: add your logic here and delete this line
	pbContactReq := new(pb.UpdateUserContactReq)
	err = copier.Copy(pbContactReq, req)
	if err != nil {
		return nil, err
	}

	content, err := json.Marshal(req.Content)
	if err != nil {
		return nil, err
	}
	pbContactReq.Content = string(content)

	contact, err := l.svcCtx.UsercenterRpc.UpdateUserContact(l.ctx, pbContactReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("修改联系方式失败"), "add Contact rpc AddUserContact fail req: %+v , err : %v ", req, err)
	}
	return &types.UpDateContactResp{Id: contact.Id}, nil

}
