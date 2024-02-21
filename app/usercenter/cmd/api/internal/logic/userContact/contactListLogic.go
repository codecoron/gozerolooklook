package userContact

import (
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ContactListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewContactListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ContactListLogic {
	return &ContactListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ContactListLogic) ContactList(req *types.ContactListReq) (resp *types.ContactListResp, err error) {
	rpcContactList, err := l.svcCtx.UsercenterRpc.SearchUserContact(l.ctx, &usercenter.SearchUserContactReq{
		Page:   req.Page,
		Limit:  req.PageSize,
		UserId: ctxdata.GetUidFromCtx(l.ctx), //只取自己的联系方式
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get SearchLottery"), "Failed to get SearchLottery err : %v ,req:%+v", err, req)
	}

	var ContactList []types.Contact
	if len(rpcContactList.UserContact) > 0 {
		for _, item := range rpcContactList.UserContact {
			var t types.Contact
			_ = copier.Copy(&t, item)
			_ = json.Unmarshal([]byte(item.Content), &t.Content)
			ContactList = append(ContactList, t)
		}
	}
	return &types.ContactListResp{List: ContactList}, nil
}
