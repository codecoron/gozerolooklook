package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/model"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserContactLogic {
	return &SearchUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserContactLogic) SearchUserContact(in *pb.SearchUserContactReq) (*pb.SearchUserContactResp, error) {
	list, err := l.svcCtx.UserContactModel.FindPageByUserId(l.ctx, in.UserId, in.Page, in.Limit)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user contact err : %v , in :%+v", err, in)
	}

	var resp []*pb.UserContact
	if len(list) > 0 {
		for _, contact := range list {
			var pbContact pb.UserContact
			_ = copier.Copy(&pbContact, contact)
			resp = append(resp, &pbContact)
		}
	}

	return &pb.SearchUserContactResp{
		UserContact: resp,
	}, nil
}
