package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"
	"looklook/app/notice/model"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGetNoticeSubscribePreferenceFail = xerr.NewErrMsg("get notice subscribe preference fail")

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
	subscribePreference, err := l.svcCtx.NoticeSubscribePreferenceModel.FindOneByUserOpenidMsgTemplateId(l.ctx, in.Openid, in.TemplateId)
	if errors.Is(err, model.ErrNotFound) {
		return &pb.GetNoticeSubscribePreferenceResp{}, nil
	}
	if err != nil {
		return nil, errors.Wrapf(ErrGetNoticeSubscribePreferenceFail, "Failed to query the preference, NoticeSubscribePreferenceModel FindOneByUserOpenidMsgTemplateId fail , req : %+v , err : %v", in, err)
	}

	return &pb.GetNoticeSubscribePreferenceResp{
		Id:          subscribePreference.Id,
		Openid:      subscribePreference.UserOpenid,
		TemplateId:  subscribePreference.MsgTemplateId,
		AcceptCount: subscribePreference.AcceptCount,
	}, nil
}
