package userDynamic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/xerr"

	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDynamicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDynamicLogic {
	return &CreateDynamicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDynamicLogic) CreateDynamic(req *types.CreateDynamicReq) (resp *types.CreateDynamicResp, err error) {
	pbDynamicReq := new(pb.AddUserDynamicReq)
	err = copier.Copy(pbDynamicReq, req)
	if err != nil {
		return nil, err
	}
	if req.UserId == 0 {
		return nil, errors.New("用户ID不能为空")
	}
	if pbDynamicReq.DynamicUrl == "" {
		return nil, errors.New("图片不能为空")
	}
	if pbDynamicReq.Remark == "" {
		return nil, errors.New("备注不能为空")
	}
	addDynamic, err := l.svcCtx.UsercenterRpc.AddUserDynamic(l.ctx, pbDynamicReq)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("add dynamic fail"), "add dynamic rpc AddUserAddress fail req: %+v , err : %v ", req, err)
	}

	return &types.CreateDynamicResp{Id: addDynamic.Id}, nil

}
