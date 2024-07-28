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
	"looklook/app/comment/cmd/rpc/comment"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/common/xerr"
)

var (
	dayCount  int64
	weekCount int64
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

type TaskStrategy interface {
	Run() error
}

// NewbieTaskStrategy 新手任务
type NewbieTaskStrategy struct {
	L        *GetTaskProgressLogic
	Logic    *AddTaskRecordLogic
	TaskId   int64
	UserId   int64
	Seesion  sqlx.Session
	Progress *model.TaskProgress
}

// OtherTaskStrategy 其他任务
type OtherTaskStrategy struct {
	L       *GetTaskProgressLogic
	Logic   *AddTaskRecordLogic
	TaskId  int64
	UserId  int64
	Seesion sqlx.Session
}

func (l *GetTaskProgressLogic) GetTaskProgress(in *pb.GetTaskProgressReq) (*pb.GetTaskProgressResp, error) {
	progress := &model.TaskProgress{}
	out := pb.GetTaskProgressResp{}
	logic := NewAddTaskRecordLogic(l.ctx, l.svcCtx)
	var strategy TaskStrategy
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

	// 现在progress肯定有内容了，开始查询任务完成进度
	err = l.svcCtx.TaskProgressModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		// 新手任务（首先查询表task_record中有无数据，如果有说明完成）
		// 任务一：参与任意抽奖
		if progress.IsParticipatedLottery != 1 {
			strategy = &NewbieTaskStrategy{
				L:        l,
				Logic:    logic,
				TaskId:   1,
				UserId:   in.UserId,
				Seesion:  session,
				Progress: progress,
			}
			err := strategy.Run()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "NewbieTaskStrategy 1 run error: %v", err)
			}
		}
		// 任务二：订阅签到提醒（UpdateSub中完成）
		// 任务三：发起任意抽奖
		if progress.IsInitiatedLottery != 1 {
			strategy = &NewbieTaskStrategy{
				L:        l,
				Logic:    logic,
				TaskId:   3,
				UserId:   in.UserId,
				Seesion:  session,
				Progress: progress,
			}
			err := strategy.Run()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "NewbieTaskStrategy 3 run error: %v", err)
			}
		}

		// 每日任务（首先查询首先查询表task_record中今天之内有无数据，如果有说明完成）
		// 任务四：参加3个首页抽奖
		_, err := l.svcCtx.TaskRecordModel.FindByUserIdAndTaskIdByDay(l.ctx, in.UserId, 4)
		if err == sqlc.ErrNotFound {
			strategy = &OtherTaskStrategy{
				L:       l,
				Logic:   logic,
				TaskId:  4,
				UserId:  in.UserId,
				Seesion: session,
			}
			err := strategy.Run()
			out.DayCount = dayCount
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "NewbieTaskStrategy 4 run error: %v", err)
			}
		} else if err != nil {
			// 其他错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord 4, err: %v", err)
		} else {
			// 今日已完成
			out.DayCount = 3
		}
		// todo 任务五：观看完整视频1次
		// 任务六：发起抽奖并超过5个人参加
		_, err = l.svcCtx.TaskRecordModel.FindByUserIdAndTaskIdByDay(l.ctx, in.UserId, 6)
		if err == sqlc.ErrNotFound {
			// 没查询到任何一条数据，判断用户今天是否完成
			strategy = &OtherTaskStrategy{
				L:       l,
				Logic:   logic,
				TaskId:  6,
				UserId:  in.UserId,
				Seesion: session,
			}
			err := strategy.Run()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "NewbieTaskStrategy 6 run error: %v", err)
			}
		} else if err != nil {
			// 其他错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord 6, err: %v", err)
		}

		// 每周任务（首先查询首先查询表task_record中有无数据，如果有并且是本周说明完成，不是本周的话就需要判断）
		// 任务七：参与首页抽奖30次以上
		_, err = l.svcCtx.TaskRecordModel.FindByUserIdAndTaskIdByWeek(l.ctx, in.UserId, 7)
		if err == sqlc.ErrNotFound {
			// 没查询到任何一条数据，判断用户本周是否完成
			strategy = &OtherTaskStrategy{
				L:       l,
				Logic:   logic,
				TaskId:  7,
				UserId:  in.UserId,
				Seesion: session,
			}
			err := strategy.Run()
			out.WeekCount = weekCount
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "NewbieTaskStrategy 7 run error: %v", err)
			}
		} else if err != nil {
			// 其他错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord 7, err: %v", err)
		} else {
			// 有记录，完成了,直接返回需要的数
			out.WeekCount = 30
		}

		// 任务八：发起抽奖并超过10人参与
		_, err = l.svcCtx.TaskRecordModel.FindByUserIdAndTaskIdByWeek(l.ctx, in.UserId, 8)
		if err == sqlc.ErrNotFound {
			// 没查询到任何一条数据，判断用户本周是否完成
			strategy = &OtherTaskStrategy{
				L:       l,
				Logic:   logic,
				TaskId:  8,
				UserId:  in.UserId,
				Seesion: session,
			}
			err := strategy.Run()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "NewbieTaskStrategy 6 run error: %v", err)
			}
		} else if err != nil {
			// 其他错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord 8, err: %v", err)
		}
		// 任务九：给晒单的锦鲤点个赞
		_, err = l.svcCtx.TaskRecordModel.FindByUserIdAndTaskIdByWeek(l.ctx, in.UserId, 9)
		if err == sqlc.ErrNotFound {
			// 没查询到任何一条数据，判断用户本周是否完成
			strategy = &OtherTaskStrategy{
				L:       l,
				Logic:   logic,
				TaskId:  9,
				UserId:  in.UserId,
				Seesion: session,
			}
			err := strategy.Run()
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "NewbieTaskStrategy 6 run error: %v", err)
			}
		} else if err != nil {
			// 其他错误
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord 9, err: %v", err)
		}
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

func (s *NewbieTaskStrategy) Run() error {
	_, err := s.L.svcCtx.TaskRecordModel.FindByUserIdAndTaskId(s.L.ctx, s.UserId, s.TaskId)
	if err == sqlc.ErrNotFound {
		// 没查询到数据，判断用户是否完成
		if s.TaskId == 1 {
			check, err := s.L.svcCtx.LotteryRpc.CheckSelectedLotteryParticipated(s.L.ctx, &lottery.CheckSelectedLotteryParticipatedReq{
				UserId: s.UserId,
			})
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to CheckUserCreatedLottery, err: %v", err)
			}
			// 如果返回1，说明用户已完成该任务，增加任务记录，返回0不做处理
			if check.Participated == 1 {
				addTaskRecord := &pb.AddTaskRecordReq{
					UserId: s.UserId,
					TaskId: s.TaskId,
				}
				_, err := s.Logic.AddTaskRecord(addTaskRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 1, err: %v", err)
				}
				// 修改task_progress记录
				s.Progress.IsParticipatedLottery = 1
				err = s.L.svcCtx.TaskProgressModel.TransUpdateByUserId(s.L.ctx, s.Seesion, s.Progress)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update progress.IsParticipatedLottery, err: %v", err)
				}
			}
		} else if s.TaskId == 3 {
			check, err := s.L.svcCtx.LotteryRpc.CheckUserCreatedLottery(s.L.ctx, &lottery.CheckUserCreatedLotteryReq{
				UserId: s.UserId,
			})
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to CheckUserCreatedLottery, err: %v", err)
			}
			// 如果返回1，说明用户已完成该任务，增加任务记录，返回0不做处理
			if check.IsCreated == 1 {
				addTaskRecord := &pb.AddTaskRecordReq{
					UserId: s.UserId,
					TaskId: s.TaskId,
				}
				_, err := s.Logic.AddTaskRecord(addTaskRecord)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 3, err: %v", err)
				}
				// 修改task_progress记录
				s.Progress.IsInitiatedLottery = 1
				err = s.L.svcCtx.TaskProgressModel.TransUpdateByUserId(s.L.ctx, s.Seesion, s.Progress)
				if err != nil {
					return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to update progress.IsInitiatedLottery, err: %v", err)
				}
			}
		} else {
			return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "TaskId %d not found", s.TaskId)
		}
	} else if err != nil {
		// 其他错误
		return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to find taskRecord %d, err: %v", s.TaskId, err)
	}
	return nil
}

func (s *OtherTaskStrategy) Run() error {
	switch s.TaskId {
	case 4:
		check, err := s.L.svcCtx.LotteryRpc.GetSelectedLotteryStatistic(s.L.ctx, &lottery.GetSelectedLotteryStatisticReq{
			UserId: s.UserId,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to GetSelectedLotteryStatistic, err: %v", err)
		}
		if check.DayCount >= 3 {
			addTaskRecord := &pb.AddTaskRecordReq{
				UserId: s.UserId,
				TaskId: 4,
			}
			_, err := s.Logic.AddTaskRecord(addTaskRecord)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 4, err: %v", err)
			}
			dayCount = 3
		} else {
			// 返回今天参加首页抽奖的数量
			dayCount = check.DayCount
		}
	case 6:
		check, err := s.L.svcCtx.LotteryRpc.CheckUserCreatedLotteryAndToday(s.L.ctx, &lottery.CheckUserCreatedLotteryAndTodayReq{
			UserId: s.UserId,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to CheckUserCreatedLotteryAndToday, err: %v", err)
		}
		if check.Yes == 1 {
			addTaskRecord := &pb.AddTaskRecordReq{
				UserId: s.UserId,
				TaskId: 6,
			}
			_, err := s.Logic.AddTaskRecord(addTaskRecord)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 6, err: %v", err)
			}
		}
	case 7:
		check, err := s.L.svcCtx.LotteryRpc.GetSelectedLotteryStatistic(s.L.ctx, &lottery.GetSelectedLotteryStatisticReq{
			UserId: s.UserId,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to GetSelectedLotteryStatistic, err: %v", err)
		}
		if check.WeekCount >= 30 {
			addTaskRecord := &pb.AddTaskRecordReq{
				UserId: s.UserId,
				TaskId: 7,
			}
			_, err := s.Logic.AddTaskRecord(addTaskRecord)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 7, err: %v", err)
			}
			weekCount = 30
		} else {
			// 返回本周参加首页抽奖的数量
			weekCount = check.WeekCount
		}
	case 8:
		check, err := s.L.svcCtx.LotteryRpc.CheckUserCreatedLotteryAndThisWeek(s.L.ctx, &lottery.CheckUserCreatedLotteryAndThisWeekReq{
			UserId: s.UserId,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to CheckUserCreatedLotteryAndThisWeek, err: %v", err)
		}
		if check.Yes == 1 {
			addTaskRecord := &pb.AddTaskRecordReq{
				UserId: s.UserId,
				TaskId: 8,
			}
			_, err := s.Logic.AddTaskRecord(addTaskRecord)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 8, err: %v", err)
			}
		}
	case 9:
		check, err := s.L.svcCtx.CommentRpc.CheckUserPraise(s.L.ctx, &comment.CheckUserPraiseReq{
			UserId: s.UserId,
		})
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to CheckUserPraise, err: %v", err)
		}
		if check.IsPraise == 1 {
			addTaskRecord := &pb.AddTaskRecordReq{
				UserId: s.UserId,
				TaskId: 9,
			}
			_, err := s.Logic.AddTaskRecord(addTaskRecord)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to AddTaskRecord 9, err: %v", err)
			}
		}
	default:
		return errors.Wrapf(xerr.NewErrCode(xerr.GET_TASK_PROGRESS_ERROR), "TaskId %d not found", s.TaskId)
	}
	return nil
}
