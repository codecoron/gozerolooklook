package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/notice/model"
	"looklook/common/constants"
	"looklook/common/xerr"

	"looklook/app/notice/cmd/rpc/internal/svc"
	"looklook/app/notice/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

var ErrSaveNoticeSubscribePreferenceFail = xerr.NewErrMsg("save notice subscribe preference fail")

type SaveNoticeSubscribePreferenceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveNoticeSubscribePreferenceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveNoticeSubscribePreferenceLogic {
	return &SaveNoticeSubscribePreferenceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SaveNoticeSubscribePreferenceLogic) SaveNoticeSubscribePreference(in *pb.SaveNoticeSubscribePreferenceReq) (*pb.SaveNoticeSubscribePreferenceResp, error) {
	subscribePreference, err := l.svcCtx.NoticeSubscribePreferenceModel.FindOneByUserOpenidMsgTemplateId(l.ctx, in.Openid, in.TemplateId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(ErrSaveNoticeSubscribePreferenceFail, "Failed to query the preference, NoticeSubscribePreferenceModel FindOneByUserOpenidMsgTemplateId fail , req : %+v , err : %v", in, err)
	}
	if subscribePreference == nil {
		subscribePreference = &model.NoticeSubscribePreference{
			UserOpenid:    in.Openid,
			MsgTemplateId: in.TemplateId,
		}
	}

	acceptCount := subscribePreference.AcceptCount

	switch in.Type {
	case constants.TypeAcceptSubscribeMessage:
		acceptCount++
	case constants.TypeRejectSubscribeMessage:
		acceptCount = 0
	case constants.TypeSendSubscribeMessage:
		if acceptCount > 0 {
			acceptCount--
		} else {
			return nil, errors.Wrapf(ErrSaveNoticeSubscribePreferenceFail, "SaveNoticeSubscribePreferenceLogic insufficient accept count, req : %+v", in)
		}
	default:
		return nil, errors.Wrapf(ErrSaveNoticeSubscribePreferenceFail, "SaveNoticeSubscribePreferenceLogic invalid type, req : %+v", in)
	}

	subscribePreference.AcceptCount = acceptCount

	_, err = l.svcCtx.NoticeSubscribePreferenceModel.Upsert(l.ctx, subscribePreference)
	if err != nil {
		return nil, errors.Wrapf(ErrSaveNoticeSubscribePreferenceFail, "Failed to upsert the preference, NoticeSubscribePreferenceModel Upsert fail , subscribePreference : %+v , err : %v", subscribePreference, err)
	}

	return &pb.SaveNoticeSubscribePreferenceResp{}, nil
}
