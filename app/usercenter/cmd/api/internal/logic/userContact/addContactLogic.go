package userContact

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddContactLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddContactLogic {
	return &AddContactLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddContactLogic) AddContact(req *types.CreateContactReq) (resp *types.CreateContactResp, err error) {
	pbContactReq := new(pb.AddUserContactReq)
	err = copier.Copy(pbContactReq, req)
	if err != nil {
		return nil, err
	}
	pbContactReq.UserId = ctxdata.GetUidFromCtx(l.ctx)
	ContentByte, err := json.Marshal(req.Content)
	if err != nil {
		return nil, err
	}
	pbContactReq.Content = string(ContentByte)
	contact, err := l.svcCtx.UsercenterRpc.AddUserContact(l.ctx, pbContactReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("add Contact fail"), "add Contact rpc AddUserContact fail req: %+v , err : %v ", req, err)
	}
	return &types.CreateContactResp{Id: contact.Id}, nil
}
