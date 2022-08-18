package profession

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysProfessionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysProfessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysProfessionListLogic {
	return &GetSysProfessionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysProfessionListLogic) GetSysProfessionList() (resp *types.SysProfessionListResp, err error) {
	sysProfessionList, err := l.svcCtx.SysProfessionModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var profession types.Profession
	var professionList []types.Profession
	for _, v := range sysProfessionList {
		err := copier.Copy(&profession, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		professionList = append(professionList, profession)
	}

	return &types.SysProfessionListResp{
		ProfessionList: professionList,
	}, nil
}
