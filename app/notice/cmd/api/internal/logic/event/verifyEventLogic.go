package event

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	"looklook/common/xerr"
	"net/http"
	"sort"

	"looklook/app/notice/cmd/api/internal/svc"
	"looklook/app/notice/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

// ErrWxMiniAuthFailError error
var ErrWxEventSignatureInvalidError = xerr.NewErrMsg("wechat event signature is invalid")

type VerifyEventLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyEventLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyEventLogic {
	return &VerifyEventLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyEventLogic) VerifyEvent(req *types.VerifyEventReq, w http.ResponseWriter) (resp *types.VerifyEventResp, err error) {
	// 参考链接：https://developers.weixin.qq.com/miniprogram/dev/framework/server-ability/message-push.html
	token := l.svcCtx.Config.WxMsgConf.EventToken
	tmpSlice := []string{token, req.Timestamp, req.Nonce}
	sort.Strings(tmpSlice)
	tmpStr := ""
	for _, v := range tmpSlice {
		tmpStr += v
	}

	h := sha1.New()
	h.Write([]byte(tmpStr))
	sha1Str := hex.EncodeToString(h.Sum([]byte("")))
	if sha1Str != req.Signature {
		return nil, errors.Wrapf(ErrWxEventSignatureInvalidError, "Verify event err : %v ,req:%+v", err, req)
	}

	// 特殊处理，按照回调要求返回响应，不需要返回json
	w.Header().Set(httpx.ContentType, "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(req.Echostr))
	if err != nil {
		return nil, errors.Wrapf(err, "Write event resp error : %v ,req:%+v", err, req)
	}
	return
}
