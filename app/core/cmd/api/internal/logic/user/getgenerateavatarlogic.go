package user

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGenerateAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGenerateAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGenerateAvatarLogic {
	return &GetGenerateAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGenerateAvatarLogic) GetGenerateAvatar() (resp *types.GenerateAvatarResp, err error) {
	return &types.GenerateAvatarResp{
		AvatarUrl: utils.AvatarUrl(),
	}, nil
}
