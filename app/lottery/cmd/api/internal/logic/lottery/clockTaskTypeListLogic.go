package lottery

import (
	"context"
	"fmt"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/common/constants"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClockTaskTypeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClockTaskTypeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClockTaskTypeListLogic {
	return &ClockTaskTypeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClockTaskTypeListLogic) ClockTaskTypeList(req *types.ClockTaskTypeListReq) (resp *types.ClockTaskTypeListResp, err error) {
	var ClockTypeList []types.CockTaskType
	// 构造数据
	clockTaskType := types.CockTaskType{}
	clockTaskType.Type = constants.ExperienceMiniPrograms
	clockTaskType.Seconds = constants.ExperienceMiniProgramsSeconds
	clockTaskType.Text = fmt.Sprintf(constants.ExperienceMiniProgramsText, constants.ExperienceMiniProgramsSeconds)
	ClockTypeList = append(ClockTypeList, clockTaskType)

	clockTaskType.Type = constants.BrowseOfficialAccountArticles
	clockTaskType.Seconds = constants.BrowseOfficialAccountArticlesSeconds
	clockTaskType.Text = fmt.Sprintf(constants.BrowseOfficialAccountArticlesText, constants.BrowseOfficialAccountArticlesSeconds)
	ClockTypeList = append(ClockTypeList, clockTaskType)

	clockTaskType.Type = constants.BrowseImage
	clockTaskType.Seconds = constants.BrowseImageSeconds
	clockTaskType.Text = fmt.Sprintf(constants.BrowseImageText, constants.BrowseImageSeconds)
	ClockTypeList = append(ClockTypeList, clockTaskType)

	clockTaskType.Type = constants.BrowseVideo
	clockTaskType.Seconds = constants.BrowseVideoSeconds
	clockTaskType.Text = fmt.Sprintf(constants.BrowseVideoText, constants.BrowseVideoSeconds)
	ClockTypeList = append(ClockTypeList, clockTaskType)

	return &types.ClockTaskTypeListResp{List: ClockTypeList}, nil
}
