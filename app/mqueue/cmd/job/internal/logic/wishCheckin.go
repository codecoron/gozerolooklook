package logic

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/pkg/errors"
	"looklook/app/checkin/cmd/rpc/checkin"
	"looklook/app/checkin/cmd/rpc/pb"
	"looklook/app/mqueue/cmd/job/internal/svc"
	"looklook/app/notice/cmd/rpc/notice"
	"looklook/common/xerr"
)

var WishCheckinHandlerFail = xerr.NewErrMsg("WishCheckinHandler ProcessTask fail")

type WishCheckinHandler struct {
	svcCtx *svc.ServiceContext
}

func NewWishCheckinHandler(svcCtx *svc.ServiceContext) *WishCheckinHandler {
	return &WishCheckinHandler{
		svcCtx: svcCtx,
	}
}
func (l *WishCheckinHandler) ProcessTask(ctx context.Context, _ *asynq.Task) error {
	//调用心愿rpc获取数据(用户、奖励、累计天数)
	checkinResp, err := l.svcCtx.CheckInRpc.NoticeWishCheckin(ctx, &checkin.NoticeWishCheckinReq{})
	if err != nil {
		return errors.Wrapf(WishCheckinHandlerFail, "CheckInRpc fail:%v", err)
	}
	//循环发送
	//用户数较大的情况可能存在并发性能问题，考虑是否进行mq消峰
	for _, v := range checkinResp.WishCheckinDatas {
		go func(v *pb.NoticeWishCheckinData) {
			_, err := l.svcCtx.NoticeRpc.NoticeWishSign(ctx, &notice.NoticeWishSignInReq{
				UserId:     v.UserId,
				Reward:     v.Reward,
				Accumulate: v.Accumulate,
			})
			if err != nil {
				errors.Wrapf(WishCheckinHandlerFail, "user:%d send message fail:%v", v.UserId, err)
			}
		}(v)
	}
	return nil
}
