package dictionary

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetParamDictionarySetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetParamDictionarySetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetParamDictionarySetLogic {
	return &GetParamDictionarySetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetParamDictionarySetLogic) GetParamDictionarySet() (resp *types.ParamDictionarySetResp, err error) {
	paramDictionaryList, err := l.svcCtx.SysDictionaryModel.FindDictionarySet(l.ctx)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var dictionary types.ParamDictionary
	dictionaryList := make([]types.ParamDictionary, 0)
	for _, v := range paramDictionaryList {
		err := copier.Copy(&dictionary, &v)
		if err != nil {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		dictionaryList = append(dictionaryList, dictionary)
	}

	return &types.ParamDictionarySetResp{
		DictionaryList: dictionaryList,
	}, nil
}
