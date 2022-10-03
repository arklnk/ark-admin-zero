package job

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysJobPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysJobPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysJobPageLogic {
	return &GetSysJobPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysJobPageLogic) GetSysJobPage(req *types.SysJobPageReq) (resp *types.SysJobPageResp, err error) {
	sysJobList, err := l.svcCtx.SysJobModel.FindPage(l.ctx, req.Page, req.Limit)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var job types.Job
	jobList := make([]types.Job, 0)
	for _, sysJob := range sysJobList {
		err := copier.Copy(&job, &sysJob)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
		jobList = append(jobList, job)
	}

	total, err := l.svcCtx.SysJobModel.FindCount(l.ctx)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	pagination := types.Pagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.SysJobPageResp{
		List:       jobList,
		Pagination: pagination,
	}, nil
}
