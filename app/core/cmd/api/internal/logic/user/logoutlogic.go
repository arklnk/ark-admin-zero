package user

import (
	"context"
	"strconv"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/common/errorx"
	"ark-zero-admin/common/globalkey"
	"ark-zero-admin/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout() error {
	userId := utils.GetUserId(l.ctx)
	_, err := l.svcCtx.Redis.Del(globalkey.CachePermMenuKey + strconv.FormatInt(userId, 10))
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}
	return nil
}
