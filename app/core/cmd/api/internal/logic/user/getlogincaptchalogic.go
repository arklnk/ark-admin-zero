package user

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/globalkey"
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
	captcha := utils.Captcha{
		Height: 40,
		Width:  80,
		Length: 4,
		ColorR: 40,
		ColorG: 30,
		ColorB: 89,
		ColorA: 0,
	}

	driver := captcha.DriverString()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	val := store.Get(id, true)
	captchaId := uuid.NewV4().String()
	err = l.svcCtx.Redis.Setex(globalkey.CacheLoginCaptchaKey +captchaId, val, 300)

	return &types.LoginCaptchaResp{
		CaptchaId:  captchaId,
		VerifyCode: b64s,
	}, nil
}