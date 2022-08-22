package job

import (
	"ark-admin-zero/common/errorx"
	"context"
	"github.com/jinzhu/copier"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysJobAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysJobAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysJobAllLogic {
	return &GetSysJobAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysJobAllLogic) GetSysJobAll() (resp *types.SysJobAllResp, err error) {
	sysJobList, err := l.svcCtx.SysJobModel.FindAll(l.ctx)
	var job types.Job
	jobList := make([]types.Job, 0)
	for _, sysJob := range sysJobList {
		err := copier.Copy(&job, &sysJob)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		jobList = append(jobList, job)
	}

	return &types.SysJobAllResp{
		JobList: jobList,
	}, nil
}
