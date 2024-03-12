package user

import (
	"context"
	"looklook/app/usercenter/cmd/api/internal/svc"
	"looklook/app/usercenter/cmd/api/internal/types"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	usercenterModel "looklook/app/usercenter/model"
	"looklook/common/tool"
	"looklook/common/xerr"
	"strings"

	"github.com/pkg/errors"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
)

// ErrWxMiniAuthFailError error
var ErrWxMiniAuthFailError = xerr.NewErrMsg("wechat mini auth fail")

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) WxMiniAuthLogic {
	return WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Wechat-Mini auth
func (l *WxMiniAuthLogic) WxMiniAuth(req types.WXMiniAuthReq) (*types.WXMiniAuthResp, error) {
	//1、Wechat-Mini
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.AppId,
		AppSecret: l.svcCtx.Config.WxMiniConf.Secret,
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		logx.Error("微信登录报错\n", err)
		return nil, errors.Wrapf(xerr.NewErrMsg("发起授权请求失败"), "发起授权请求失败 err : %v , code : %s  , authResult : %+v", err, req.Code, authResult)
	}
	//2、Parsing WeChat-Mini return data
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("解析数据失败"), "解析数据失败 req : %+v , err: %v , authResult:%+v ", req, err, authResult)
	}

	//3、bind user or login.
	var userId int64
	rpcRsp, err := l.svcCtx.UsercenterRpc.GetUserAuthByAuthKey(l.ctx, &usercenter.GetUserAuthByAuthKeyReq{
		AuthType: usercenterModel.UserAuthTypeSmallWX,
		AuthKey:  authResult.OpenID,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("rpc call userAuthByAuthKey err"), "rpc call userAuthByAuthKey err : %v , authResult : %+v", err, authResult)
	}
	if rpcRsp.UserAuth == nil || rpcRsp.UserAuth.Id == 0 {
		//bind user.

		//Wechat-Mini Decrypted data
		if len(req.Nickname) == 0 {
			nicknameArr := []string{userData.NickName, tool.Krand(6, tool.KC_RAND_KIND_NUM)}
			nickName := strings.Join(nicknameArr, "")
			req.Nickname = nickName
		}
		if len(req.Avatar) == 0 {
			req.Avatar = userData.AvatarURL
		}

		openId := authResult.OpenID
		//mobile := openId[len(openId)-11:] //TODO 优化逻辑
		wxMiniRegisterRsp, err := l.svcCtx.UsercenterRpc.WxMiniRegister(l.ctx, &usercenter.WXMiniRegisterReq{
			AuthKey:  openId,
			AuthType: usercenterModel.UserAuthTypeSmallWX,
			//Mobile:   mobile,
			Nickname: req.Nickname,
			Avatar:   req.Avatar,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("UsercenterRpc.Register err"), "UsercenterRpc.Register err :%v, authResult : %+v", err, authResult)
		}

		return &types.WXMiniAuthResp{
			AccessToken:  wxMiniRegisterRsp.AccessToken,
			AccessExpire: wxMiniRegisterRsp.AccessExpire,
			RefreshAfter: wxMiniRegisterRsp.RefreshAfter,
		}, nil

	} else {
		userId = rpcRsp.UserAuth.UserId
		tokenResp, err := l.svcCtx.UsercenterRpc.GenerateToken(l.ctx, &usercenter.GenerateTokenReq{
			UserId: userId,
		})
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("usercenterRpc.GenerateToken err"), "usercenterRpc.GenerateToken err :%v, userId : %d", err, userId)
		}
		return &types.WXMiniAuthResp{
			AccessToken:  tokenResp.AccessToken,
			AccessExpire: tokenResp.AccessExpire,
			RefreshAfter: tokenResp.RefreshAfter,
		}, nil
	}
}
