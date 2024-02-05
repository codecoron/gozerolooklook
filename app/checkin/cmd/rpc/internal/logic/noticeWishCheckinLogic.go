package logic

import (
	"context"
	"github.com/pkg/errors"
	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeWishCheckinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNoticeWishCheckinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeWishCheckinLogic {
	return &NoticeWishCheckinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------others----------------------
func (l *NoticeWishCheckinLogic) NoticeWishCheckin(in *pb.NoticeWishCheckinReq) (*pb.NoticeWishCheckinResp, error) {
	userIds, err := l.svcCtx.TaskProgressModel.FindAllSubId(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskProgress data, err: %v", err)
	}
	var list []*pb.NoticeWishCheckinData
	logic := NewGetCheckinRecordByUserIdLogic(l.ctx, l.svcCtx)
	for _, id := range userIds {
		taskRecord := &pb.GetCheckinRecordByUserIdReq{
			UserId: id,
		}
		resp, err := logic.GetCheckinRecordByUserId(taskRecord)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to GetTaskRecordByUserId, err: %v", err)
		}
		// 用户在提醒前已经签到，不用提醒了
		if resp.State == 1 {
			continue
		}
		var accumulate int64
		var i int64
		if resp.ContinuousCheckinDays >= 7 {
			accumulate = 1
		} else {
			accumulate = resp.ContinuousCheckinDays + 1
		}
		switch accumulate {
		case 1, 2:
			i = 5
		case 3:
			i = 10
		case 4:
			i = 15
		case 5:
			i = 20
		case 6:
			i = 30
		case 7:
			i = 40
		default:
			i = 0
		}
		userRecord := &pb.NoticeWishCheckinData{
			UserId:     id,
			Accumulate: accumulate,
			Reward:     i,
		}
		list = append(list, userRecord)
	}
	return &pb.NoticeWishCheckinResp{
		WishCheckinDatas: list,
	}, nil
}
