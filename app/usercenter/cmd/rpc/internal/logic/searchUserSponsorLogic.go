package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/model"
)

type SearchUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserSponsorLogic {
	return &SearchUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserSponsorLogic) SearchUserSponsor(in *pb.SearchUserSponsorReq) (*pb.SearchUserSponsorResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.UserSponsorModel.FindPageByUserId(l.ctx, in.UserId, in.Page, in.Limit)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user sponsor err : %v , in :%+v", err, in)
	}

	var resp []*pb.UserSponsor
	if len(list) > 0 {
		for _, sponsor := range list {
			var pbSponsor pb.UserSponsor
			_ = copier.Copy(&pbSponsor, sponsor)
			resp = append(resp, &pbSponsor)
		}
	}
	return &pb.SearchUserSponsorResp{}, nil
}
