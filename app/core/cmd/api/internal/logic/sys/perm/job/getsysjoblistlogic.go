package job

import (
	"context"

	"ark-zero-admin/app/core/cmd/api/internal/svc"
	"ark-zero-admin/app/core/cmd/api/internal/types"
	"ark-zero-admin/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysJobListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysJobListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysJobListLogic {
	return &GetSysJobListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysJobListLogic) GetSysJobList() (resp *types.SysJobListResp, err error) {
	sysJobList, err := l.svcCtx.SysJobModel.FindAll(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var job types.Job
	var jobList []types.Job
	for _, sysJob := range sysJobList {
		err := copier.Copy(&job, &sysJob)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		jobList = append(jobList, job)
	}

	return &types.SysJobListResp{
		JobList: jobList,
	}, nil
}
