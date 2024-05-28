package logic

import (
	"context"

	"microservices-go-zero/app/usercenter/cmd/rpc/internal/svc"
	"microservices-go-zero/app/usercenter/cmd/rpc/usercenter"
	"microservices-go-zero/app/usercenter/model"
	"microservices-go-zero/common/tool"
	"microservices-go-zero/common/xerr"
	
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

var ErrGenerateTokenError = xerr.NewErrMsg("Failed to generate token")
var ErrUsernamePwdError = xerr.NewErrMsg("The account or password is incorrect")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *usercenter.LoginReq) (*usercenter.LoginResp, error) {
	var userId int64
	var err error
	// to-do: wx mini program auth
	switch in.AuthType {
		case model.UserAuthTypeSystem:
			userId, err = l.loginByMobile(in.AuthKey, in.Password)
		default:
			return nil, xerr.NewErrCode(xerr.SERVER_COMMON_ERROR)
	}
	if err != nil {
		return nil, err
	}

	// generate the token, so that the service doesn't call rpc internally
	generateTokenLogic := NewGenerateTokenLogic(l.ctx,l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errors.Wrapf(ErrGenerateTokenError, "GenerateToken userId : %d", userId)
	}

	return &usercenter.LoginResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}

func (l *LoginLogic) loginByMobile(mobile, password string) (int64, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx,mobile)
	if err != nil && err != model.ErrNotFound {
		return 0, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Querying user information based on mobile phone number failed: mobile:%s, err:%v", mobile, err)
	}
	if user == nil {
		return 0, errors.Wrapf(ErrUserNoExistsError, "mobile:%s", mobile)
	}

	if !(tool.Md5ByString(password) == user.Password) {
		return 0, errors.Wrap(ErrUsernamePwdError, "Password matching error")
	}

	return user.Id, nil
}