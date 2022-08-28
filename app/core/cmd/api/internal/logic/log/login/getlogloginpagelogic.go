package login

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/config"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogLoginPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogLoginPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogLoginPageLogic {
	return &GetLogLoginPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogLoginPageLogic) GetLogLoginPage(req *types.LogLoginPageReq) (resp *types.LogLoginPageResp, err error) {
	loginLogList, err := l.svcCtx.SysLogModel.FindPage(l.ctx, config.SysLoginLogType, req.Page, req.Limit)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var loginLog types.LogLogin
	logList := make([]types.LogLogin, 0)
	for _, v := range loginLogList {
		err := copier.Copy(&loginLog, &v)
		loginLog.CreateTime = v.CreateTime.Format(config.SysDateFormat)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		logList = append(logList, loginLog)
	}

	total, err := l.svcCtx.SysLogModel.FindCount(l.ctx, config.SysLoginLogType)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	pagination := types.LogLoginPagePagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.LogLoginPageResp{
		LogLoginList: logList,
		Pagination:   pagination,
	}, nil
}
