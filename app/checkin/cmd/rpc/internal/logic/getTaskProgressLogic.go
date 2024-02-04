package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"
	"looklook/app/checkin/model"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/xerr"
	"time"
)

type GetTaskProgressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaskProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaskProgressLogic {
	return &GetTaskProgressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaskProgressLogic) GetTaskProgress(in *pb.GetTaskProgressReq) (*pb.GetTaskProgressResp, error) {
	// todo: 调用其他服务，查询任务进度
	progress := &model.TaskProgress{}
	out := pb.GetTaskProgressResp{}
	logic := NewAddTaskRecordLogic(l.ctx, l.svcCtx)
	getProgress, err := l.svcCtx.TaskProgressModel.FindOneByUserId(l.ctx, in.UserId)
	if err == sqlc.ErrNotFound {
		// 没查询到，新增数据
		progress.UserId = in.UserId
		insert, err := l.svcCtx.TaskProgressModel.InsertByUserId(l.ctx, progress)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to insert taskProgress data : %+v , err: %v", progress, err)
		}
		progress.Id, _ = insert.LastInsertId()
	} else if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskProgress data : %+v , err: %v", progress, err)
	} else {
		_ = copier.Copy(progress, getProgress)
	}
	logx.Error("progress:", progress)

	// 现在progress肯定有内容了，开始查询任务完成进度
	err = l.svcCtx.TaskProgressModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		// 新手任务（首先查询表task_record中有无数据，如果有说明完成）
		// 任务一：参与任意抽奖
		if progress.IsParticipatedLottery != 1 {
			_, err := l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(l.ctx, in.UserId, 1)
			if err == sqlc.ErrNotFound {
				// 没查询到数据，判断用户是否完成
				check, err := l.svcCtx.LotteryRpc.CheckSelectedLotteryParticipated(l.ctx, &lottery.CheckSelectedLotteryParticipatedReq{
					UserId: in.UserId,
				})
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to CheckUserCreatedLottery, err: %v", err)
				}
				// 如果返回1，说明用户已完成该任务，增加任务记录，返回0不做处理
				if check.Participated == 1 {
					addTaskRecord := &pb.AddTaskRecordReq{
						UserId: in.UserId,
						TaskId: 1,
					}
					_, err := logic.AddTaskRecord(addTaskRecord)
					if err != nil {
						return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 1, err: %v", err)
					}
					// 修改task_progress记录
					progress.IsParticipatedLottery = 1
					err = l.svcCtx.TaskProgressModel.TransUpdateByUserId(l.ctx, session, progress)
					if err != nil {
						return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update progress.IsParticipatedLottery, err: %v", err)
					}
				}
			} else if err != nil {
				// 其他错误
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord, err: %v", err)
			}
			// todo 任务二：订阅签到提醒

			// 任务三：发起任意抽奖
			_, err = l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(l.ctx, in.UserId, 3)
			if err == sqlc.ErrNotFound {
				// 没查询到数据，判断用户是否完成
				check, err := l.svcCtx.LotteryRpc.CheckUserCreatedLottery(l.ctx, &lottery.CheckUserCreatedLotteryReq{
					UserId: in.UserId,
				})
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to CheckUserCreatedLottery, err: %v", err)
				}
				// 如果返回1，说明用户已完成该任务，增加任务记录，返回0不做处理
				if check.IsCreated == 1 {
					addTaskRecord := &pb.AddTaskRecordReq{
						UserId: in.UserId,
						TaskId: 3,
					}
					_, err := logic.AddTaskRecord(addTaskRecord)
					if err != nil {
						return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 3, err: %v", err)
					}
					// 修改task_progress记录
					progress.IsInitiatedLottery = 1
					err = l.svcCtx.TaskProgressModel.TransUpdateByUserId(l.ctx, session, progress)
					if err != nil {
						return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update progress.IsInitiatedLottery, err: %v", err)
					}
				}
			} else if err != nil {
				// 其他错误
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord 3, err: %v", err)
			}
		}
		// 每日任务（首先查询首先查询表task_record中有无数据，如果有并且是今天说明完成，不是今天的话就需要判断）
		// 任务四：参加3个首页抽奖
		taskRecord, err := l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(l.ctx, in.UserId, 4)
		if err == sqlc.ErrNotFound {
			// 没查询到任何一条数据，判断用户今天是否完成
			check, err := l.svcCtx.LotteryRpc.GetSelectedLotteryStatistic(l.ctx, &lottery.GetSelectedLotteryStatisticReq{
				UserId: in.UserId,
			})
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to GetSelectedLotteryStatistic, err: %v", err)
			}
			logx.Error("day,week num :", check.DayCount, check.WeekCount)
			if check.DayCount >= 3 {
				addTaskRecord := &pb.AddTaskRecordReq{
					UserId: in.UserId,
					TaskId: 4,
				}
				_, err := logic.AddTaskRecord(addTaskRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 4, err: %v", err)
				}
			}
			// 返回今天参加首页抽奖的数量
			out.DayCount = check.DayCount
		} else if err != nil {
			// 其他错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord, err: %v", err)
		} else {
			// 查到数据了
			// 获取今天的日期
			today := time.Now().Format("2006-01-02")
			// 判断 taskRecord.CreateTime 是否是今天
			if taskRecord.CreateTime.Format("2006-01-02") != today {
				// 如果不是今天，判断用户今天是否完成
				check, err := l.svcCtx.LotteryRpc.GetSelectedLotteryStatistic(l.ctx, &lottery.GetSelectedLotteryStatisticReq{
					UserId: in.UserId,
				})
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to GetSelectedLotteryStatistic, err: %v", err)
				}
				if check.DayCount >= 3 {
					addTaskRecord := &pb.AddTaskRecordReq{
						UserId: in.UserId,
						TaskId: 4,
					}
					_, err := logic.AddTaskRecord(addTaskRecord)
					if err != nil {
						return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 4, err: %v", err)
					}
				}
				// 返回今天参加首页抽奖的数量
				out.DayCount = check.DayCount
			} else {
				// 有记录，完成了,直接返回需要的数
				out.DayCount = 3
			}
		}
		// todo 任务五：观看完整视频1次
		// todo 任务六：发起抽奖并超过5个人参加

		// 每周任务（首先查询首先查询表task_record中有无数据，如果有并且是本周说明完成，不是本周的话就需要判断）
		// 任务七：参与首页抽奖30次以上
		taskRecord, err = l.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(l.ctx, in.UserId, 7)
		if err == sqlc.ErrNotFound {
			// 没查询到任何一条数据，判断用户今天是否完成
			check, err := l.svcCtx.LotteryRpc.GetSelectedLotteryStatistic(l.ctx, &lottery.GetSelectedLotteryStatisticReq{
				UserId: in.UserId,
			})
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to GetSelectedLotteryStatistic, err: %v", err)
			}
			logx.Error("day,week num :", check.DayCount, check.WeekCount)
			if check.WeekCount >= 30 {
				addTaskRecord := &pb.AddTaskRecordReq{
					UserId: in.UserId,
					TaskId: 7,
				}
				_, err := logic.AddTaskRecord(addTaskRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 7, err: %v", err)
				}
			}
			// 返回本周参加首页抽奖的数量
			out.WeekCount = check.WeekCount
		} else if err != nil {
			// 其他错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord, err: %v", err)
		} else {
			// 查到数据了
			// 获取当前时间
			now := time.Now()
			// 获取本周的开始时间和结束时间
			startOfWeek := now.AddDate(0, 0, -int(now.Weekday())).Truncate(24 * time.Hour)
			endOfWeek := startOfWeek.AddDate(0, 0, 7)
			// 判断 taskRecord.CreateTime 是否是本周之内
			if !taskRecord.CreateTime.After(startOfWeek) || !taskRecord.CreateTime.Before(endOfWeek) {
				// 如果不是本周，判断用户本周是否完成
				check, err := l.svcCtx.LotteryRpc.GetSelectedLotteryStatistic(l.ctx, &lottery.GetSelectedLotteryStatisticReq{
					UserId: in.UserId,
				})
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to GetSelectedLotteryStatistic, err: %v", err)
				}
				if check.WeekCount >= 30 {
					addTaskRecord := &pb.AddTaskRecordReq{
						UserId: in.UserId,
						TaskId: 7,
					}
					_, err := logic.AddTaskRecord(addTaskRecord)
					if err != nil {
						return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 7, err: %v", err)
					}
				}
				// 返回本周参加首页抽奖的数量
				out.WeekCount = check.WeekCount
			} else {
				// 有记录，完成了,直接返回需要的数
				out.WeekCount = 30
			}
		}
		// todo 任务八：发起抽奖并超过10人参与
		// todo 任务九：给晒单的锦鲤点个赞
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &pb.GetTaskProgressResp{
		DayCount:  out.DayCount,
		WeekCount: out.WeekCount,
	}, nil
}
