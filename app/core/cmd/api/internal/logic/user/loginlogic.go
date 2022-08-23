package user

import (
	"context"
	"net/http"
	"time"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"
	"ark-admin-zero/common/utils"

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

func (l *LoginLogic) Login(req *types.LoginReq, r *http.Request) (resp *types.LoginResp, err error) {
	verifyCode, _ := l.svcCtx.Redis.Get(globalkey.SysLoginCaptchaCachePrefix + req.CaptchaId)
	if verifyCode != req.VerifyCode {
		return nil, errorx.NewDefaultError(errorx.CaptchaErrorCode)
	}

	sysUser, err := l.svcCtx.SysUserModel.FindOneByAccount(l.ctx, req.Account)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.AccountErrorCode)
	}

	if sysUser.Password != utils.MD5(req.Password+l.svcCtx.Config.Salt) {
		return nil, errorx.NewDefaultError(errorx.PasswordErrorCode)
	}

	if sysUser.Status != globalkey.SysEnable {
		return nil, errorx.NewDefaultError(errorx.AccountDisableErrorCode)
	}

	token, _ := l.getJwtToken(sysUser.Id)
	_, err = l.svcCtx.Redis.Del(req.CaptchaId)

	loginLog := model.SysLog{
		UserId:  sysUser.Id,
		Ip:      r.RemoteAddr,
		Uri:     r.RequestURI,
		Type:    1,
		Status:  1,
	}
	_, err = l.svcCtx.SysLogModel.Insert(l.ctx, &loginLog)

	return &types.LoginResp{
		Token: token,
	}, nil
}

func (l *LoginLogic) getJwtToken(userId int64) (string, error) {
	iat := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + l.svcCtx.Config.JwtAuth.AccessExpire
	claims["iat"] = iat
	claims[globalkey.SysJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
}
