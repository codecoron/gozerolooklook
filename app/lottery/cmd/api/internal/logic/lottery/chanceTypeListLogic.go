package lottery

import (
	"context"
	"fmt"
	"looklook/common/constants"

	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChanceTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChanceTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChanceTypeListLogic {
	return &ChanceTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChanceTypeListLogic) ChanceTypeList(req *types.ChanceTypeListReq) (resp *types.ChanceTypeListResp, err error) {
	var ChanceTypeList []types.ChanceType
	changeType := types.ChanceType{}
	changeType.Type = constants.Random
	changeType.Text = constants.RandomText
	ChanceTypeList = append(ChanceTypeList, changeType)

	for i := 1; i <= 10; i++ {
		changeType.Type = constants.Appoint
		changeType.Text = fmt.Sprintf(constants.AppointText, i)
		ChanceTypeList = append(ChanceTypeList, changeType)
	}
	return &types.ChanceTypeListResp{List: ChanceTypeList}, nil
}
