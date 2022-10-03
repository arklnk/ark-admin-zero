package dict

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetConfigDictListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetConfigDictListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConfigDictListLogic {
	return &GetConfigDictListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetConfigDictListLogic) GetConfigDictList() (resp *types.ConfigDictListResp, err error) {
	configDictionaryList, err := l.svcCtx.SysDictionaryModel.FindDictionaryList(l.ctx)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var dictionary types.ConfigDict
	dictionaryList := make([]types.ConfigDict, 0)
	for _, v := range configDictionaryList {
		err := copier.Copy(&dictionary, &v)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
		dictionaryList = append(dictionaryList, dictionary)
	}

	return &types.ConfigDictListResp{
		List: dictionaryList,
	}, nil
}
