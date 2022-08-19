package job

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

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

func (l *GetSysJobListLogic) GetSysJobList(req *types.SysJobListReq) (resp *types.SysJobListResp, err error) {
	sysJobList, err := l.svcCtx.SysJobModel.FindByPage(l.ctx, req.Page, req.Limit)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var job types.Job
	jobList := make([]types.Job, 0)
	for _, sysJob := range sysJobList {
		err := copier.Copy(&job, &sysJob)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		jobList = append(jobList, job)
	}

	total, err := l.svcCtx.SysJobModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	pagination := types.SysJobListPagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.SysJobListResp{
		JobList:    jobList,
		Pagination: pagination,
	}, nil
}
