package dictionary

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetParamDictionaryPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetParamDictionaryPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParamDictionaryPageLogic {
	return &GetParamDictionaryPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetParamDictionaryPageLogic) GetParamDictionaryPage(req *types.ParamDictionaryPageReq) (resp *types.ParamDictionaryPageResp, err error) {
	paramDictionaryList, err := l.svcCtx.SysDictionaryModel.FindPageByParentId(l.ctx, req.ParentId, req.Page, req.Limit)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var dictionary types.ParamDictionary
	dictionaryList := make([]types.ParamDictionary, 0)
	for _, sysDictionary := range paramDictionaryList {
		err := copier.Copy(&dictionary, &sysDictionary)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		dictionaryList = append(dictionaryList, dictionary)
	}

	total, err := l.svcCtx.SysDictionaryModel.FindCountByParentId(l.ctx, req.ParentId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	pagination := types.ParamDictionaryPagination{
		Page:  req.Page,
		Limit: req.Limit,
		Total: total,
	}

	return &types.ParamDictionaryPageResp{
		ParamDictionaryList: dictionaryList,
		Pagination:          pagination,
	}, nil
}
