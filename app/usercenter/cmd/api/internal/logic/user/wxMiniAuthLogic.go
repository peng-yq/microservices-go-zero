package user

import (
	"context"
	"fmt"

	"microservices-go-zero/app/usercenter/cmd/api/internal/svc"
	"microservices-go-zero/app/usercenter/cmd/api/internal/types"
	"microservices-go-zero/app/usercenter/cmd/rpc/usercenter"
	usercenterModel "microservices-go-zero/app/usercenter/model"
	"microservices-go-zero/common/xerr"

	"github.com/pkg/errors"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/zeromicro/go-zero/core/logx"
)

// ErrWxMiniAuthFailError error
var ErrWxMiniAuthFailError = xerr.NewErrMsg("wechat mini program auth failed")

type WxMiniAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// wechat mini program auth
func NewWxMiniAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxMiniAuthLogic {
	return &WxMiniAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxMiniAuthLogic) WxMiniAuth(req *types.WXMiniAuthReq) (resp *types.WXMiniAuthResp, err error) {
	// Wechat-Mini
	miniprogram := wechat.NewWechat().GetMiniProgram(&miniConfig.Config{
		AppID:     l.svcCtx.Config.WxMiniConf.AppId,
		AppSecret: l.svcCtx.Config.WxMiniConf.Secret,
		Cache:     cache.NewMemory(),
	})
	authResult, err := miniprogram.GetAuth().Code2Session(req.Code)
	if err != nil || authResult.ErrCode != 0 || authResult.OpenID == "" {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "Failed to initiate authorization request, err : %v , code : %s  , authResult : %+v", err, req.Code, authResult)
	}

	// Parsing WeChat-Mini return data
	userData, err := miniprogram.GetEncryptor().Decrypt(authResult.SessionKey, req.EncryptedData, req.IV)
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "Failed to parse data, req : %+v , err: %v , authResult:%+v ", req, err, authResult)
	}

	// use rpc to bind user or login
	var userId int64
	rpcRsp, err := l.svcCtx.UsercenterRpc.GetUserAuthByAuthKey(l.ctx, &usercenter.GetUserAuthByAuthKeyReq{
		AuthType: usercenterModel.UserAuthTypeSmallWX,
		AuthKey:  authResult.OpenID,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrWxMiniAuthFailError, "rpc call userAuthByAuthKey err : %v , authResult : %+v", err, authResult)
	}
	if rpcRsp.UserAuth == nil || rpcRsp.UserAuth.Id == 0 {
		//bind user.
		//Wechat-Mini Decrypted data
		mobile := userData.PhoneNumber
		nickName := fmt.Sprintf("microservices%s", mobile[7:])
		registerRsp, err := l.svcCtx.UsercenterRpc.Register(l.ctx, &usercenter.RegisterReq{
			AuthKey:  authResult.OpenID,
			AuthType: usercenterModel.UserAuthTypeSmallWX,
			Mobile:   mobile,
			Nickname: nickName,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "UsercenterRpc.Register err :%v, authResult : %+v", err, authResult)
		}

		return &types.WXMiniAuthResp{
			AccessToken:  registerRsp.AccessToken,
			AccessExpire: registerRsp.AccessExpire,
			RefreshAfter: registerRsp.RefreshAfter,
		}, nil
	} else {
		// login
		userId = rpcRsp.UserAuth.UserId
		tokenResp, err := l.svcCtx.UsercenterRpc.GenerateToken(l.ctx, &usercenter.GenerateTokenReq{
			UserId: userId,
		})
		if err != nil {
			return nil, errors.Wrapf(ErrWxMiniAuthFailError, "usercenterRpc.GenerateToken err :%v, userId : %d", err, userId)
		}
		return &types.WXMiniAuthResp{
			AccessToken:  tokenResp.AccessToken,
			AccessExpire: tokenResp.AccessExpire,
			RefreshAfter: tokenResp.RefreshAfter,
		}, nil
	}
}
