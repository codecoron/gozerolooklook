package logic

import (
	"context"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeSubscribePreferenceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNoticeSubscribePreferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeSubscribePreferenceLogic {
	return &GetNoticeSubscribePreferenceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetNoticeSubscribePreferenceLogic) GetNoticeSubscribePreference(in *pb.GetNoticeSubscribePreferenceReq) (*pb.GetNoticeSubscribePreferenceResp, error) {
	// todo: add your logic here and delete this line

	return &pb.GetNoticeSubscribePreferenceResp{}, nil
}
