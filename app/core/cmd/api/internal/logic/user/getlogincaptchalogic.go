package user

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"

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

func (l *GetLoginCaptchaLogic) GetLoginCaptcha() (resp *types.CaptchaResp, err error) {
	// todo: add your logic here and delete this line

	return
}
