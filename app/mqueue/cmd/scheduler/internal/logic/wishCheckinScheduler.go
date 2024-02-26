package logic

import (
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"looklook/app/mqueue/cmd/job/jobtype"
)

func (l *MqueueScheduler) WishCheckinScheduler() {

	task := asynq.NewTask(jobtype.ScheduleWishCheckin, nil)
	// every one minute exec
	entryID, err := l.svcCtx.Scheduler.Register("0 10 * * *", task)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("!!!MqueueSchedulerErr!!! ====> 【wishCheckinScheduler】 registered  err:%+v , task:%+v", err, task)
	}
	fmt.Printf("【wishCheckinScheduler】 registered an entry: %q \n", entryID)
}
