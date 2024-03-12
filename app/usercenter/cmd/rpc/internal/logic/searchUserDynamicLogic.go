package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"
	"looklook/app/usercenter/model"
)

type SearchUserDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserDynamicLogic {
	return &SearchUserDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserDynamicLogic) SearchUserDynamic(in *pb.SearchUserDynamicReq) (*pb.SearchUserDynamicResp, error) {
	list, err := l.svcCtx.UserDynamicModel.FindListByUserId(l.ctx, in.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	var resp []*pb.UserDynamic
	if len(list) > 0 {
		for _, dynamic := range list {
			var pbDynamic pb.UserDynamic
			_ = copier.Copy(&pbDynamic, dynamic)
			pbDynamic.Id = dynamic.Id
			pbDynamic.UserId = dynamic.UserId
			pbDynamic.DynamicUrl = dynamic.DynamicUrl
			pbDynamic.Remark = dynamic.Remark
			resp = append(resp, &pbDynamic)
		}
	}
	return &pb.SearchUserDynamicResp{
		UserDynamic: resp,
	}, nil
}
