package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/common/ctxdata"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserDynamicLogic {
	return &DelUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserDynamicLogic) DelUserDynamic(in *pb.DelUserDynamicReq) (*pb.DelUserDynamicResp, error) {

	if in.UserId == 0 || in.Id == 0 {
		return nil, errors.New("用户ID、动态ID不能为空")
	}

	if in.UserId != ctxdata.GetUidFromCtx(l.ctx) {
		return nil, errors.New("不是同一个用户ID是不能删除的")
	}
	_, err := l.svcCtx.UserDynamicModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DelUserDynamicResp{}, nil
}
