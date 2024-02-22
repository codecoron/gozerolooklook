package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"looklook/app/usercenter/model"

	"looklook/app/usercenter/cmd/rpc/internal/svc"
	"looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserSponsorLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserSponsorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserSponsorLogic {
	return &AddUserSponsorLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------抽奖发起人联系方式（抽奖赞助商）-----------------------
func (l *AddUserSponsorLogic) AddUserSponsor(in *pb.AddUserSponsorReq) (*pb.AddUserSponsorResp, error) {
	userSponsor := new(model.UserSponsor)
	err := copier.Copy(userSponsor, in)
	if err != nil {
		//todo 错误处理
		return nil, err
	}
	insert, err := l.svcCtx.UserSponsorModel.Insert(l.ctx, userSponsor)
	if err != nil {
		return nil, err
	}
	id, err := insert.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &pb.AddUserSponsorResp{Id: id}, nil
}
