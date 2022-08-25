package user

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/config"
	"ark-admin-zero/common/utils"

	"github.com/mojocn/base64Captcha"
	"github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLoginCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginCaptchaLogic {
	return &GetLoginCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLoginCaptchaLogic) GetLoginCaptcha() (resp *types.LoginCaptchaResp, err error) {
	var store = base64Captcha.DefaultMemStore
	captcha := utils.NewCaptcha(45, 80, 4, 40, 30, 89, 0)
	driver := captcha.DriverString()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	val := store.Get(id, true)
	captchaId := uuid.NewV4().String()
	err = l.svcCtx.Redis.Setex(config.SysLoginCaptchaCachePrefix+captchaId, val, 300)

	return &types.LoginCaptchaResp{
		CaptchaId:  captchaId,
		VerifyCode: b64s,
	}, nil
}
