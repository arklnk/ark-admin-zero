package dept

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysDeptListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysDeptListLogic {
	return &GetSysDeptListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysDeptListLogic) GetSysDeptList() (resp *types.SysDeptListResp, err error) {
	sysDeptList, err := l.svcCtx.SysDeptModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var dept types.Dept
	var deptList []types.Dept
	for _, sysDept := range sysDeptList {
		err := copier.Copy(&dept, &sysDept)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		deptList = append(deptList, dept)
	}

	return &types.SysDeptListResp{
		DeptList: deptList,
	}, nil
}
