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

func (l *GetSysProfessionListLogic) GetSysProfessionList(req *types.SysProfessionListReq) (resp *types.SysProfessionListResp, err error) {
	sysProfessionList, err := l.svcCtx.SysProfessionModel.FindByPage(l.ctx, req.Page, req.Limit)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var profession types.Profession
	professionList := make([]types.Profession, 0)
	for _, v := range sysProfessionList {
		err := copier.Copy(&profession, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		professionList = append(professionList, profession)
	}

	total, err := l.svcCtx.SysProfessionModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	pagination := types.Pagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.SysProfessionListResp{
		ProfessionList: professionList,
		Pagination:     pagination,
	}, nil
}
