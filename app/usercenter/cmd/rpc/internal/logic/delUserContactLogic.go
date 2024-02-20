package logic

import (
	"context"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelUserContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelUserContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelUserContactLogic {
	return &DelUserContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelUserContactLogic) DelUserContact(in *pb.DelUserContactReq) (*pb.DelUserContactResp, error) {
	//todo 做个限制 只能删除自己的  优化代码
	//l.svcCtx.UserContactModel.DeleteBatch(l.ctx,in.Id)

	for _, id := range in.Id {
		l.svcCtx.UserContactModel.Delete(l.ctx, id)
	}
	return &pb.DelUserContactResp{}, nil
}
