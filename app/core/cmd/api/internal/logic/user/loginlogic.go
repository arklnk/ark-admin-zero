package user

import (
	"ark-zero-admin/pkg/sysconstant"
	"context"
	"time"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/pkg/errorx"
	"ark-zero-admin/pkg/utils"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	verifyCode, _ := l.svcCtx.Redis.Get("captcha:" + req.CaptchaId)
	if verifyCode != req.VerifyCode {
		return nil, errorx.NewDefaultError(errorx.CaptchaErrorCode)
	}
	account, err := l.svcCtx.SysUserModel.FindOneByAccount(l.ctx, req.Account)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.AccountErrorCode)
	}
	if account.Password != utils.MD5(req.Password+l.svcCtx.Config.Salt) {
		return nil, errorx.NewDefaultError(errorx.PasswordErrorCode)
	}
	token, _ := l.getJwtToken(
		l.svcCtx.Config.JwtAuth.AccessSecret,
		time.Now().Unix(),
		l.svcCtx.Config.JwtAuth.AccessExpire,
		int64(account.Id))
	_, err = l.svcCtx.Redis.Del(req.CaptchaId)

	return &types.LoginResp{
		Token: token,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims[sysconstant.JwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
