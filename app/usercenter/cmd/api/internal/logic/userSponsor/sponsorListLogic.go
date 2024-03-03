package userSponsor

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SponsorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSponsorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SponsorListLogic {
	return &SponsorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SponsorListLogic) SponsorList(req *types.SponsorListReq) (resp *types.SponsorListResp, err error) {
	// todo: add your logic here and delete this line
	rpcSponsorList, err := l.svcCtx.UsercenterRpc.SearchUserSponsor(l.ctx, &usercenter.SearchUserSponsorReq{
		Page:   req.Page,
		Limit:  req.PageSize,
		UserId: ctxdata.GetUidFromCtx(l.ctx), //只取自己的联系方式
	})
	if err != nil {
		//todo 要使用这种写法管理错误，否则Kibana无法收集到错误日志的详情
		return nil, errors.Wrapf(xerr.NewErrMsg("Failed to get SearchLottery"), "Failed to get SearchLottery err : %v ,req:%+v", err, req)
	}

	var SponsorList []types.Sponsor
	if len(rpcSponsorList.UserSponsor) > 0 {
		for _, item := range rpcSponsorList.UserSponsor {
			var t types.Sponsor
			_ = copier.Copy(&t, item)

			SponsorList = append(SponsorList, t)
		}
	}
	return &types.SponsorListResp{List: SponsorList}, nil
}
