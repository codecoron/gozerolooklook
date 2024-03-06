package logic

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"
	"looklook/common/constants"
	"looklook/common/xerr"
	"math/rand"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClockTaskRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClockTaskRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClockTaskRecordLogic {
	return &AddClockTaskRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------完成打卡任务-----------------------
func (l *AddClockTaskRecordLogic) AddClockTaskRecord(in *pb.AddClockTaskRecordReq) (*pb.AddClockTaskRecordResp, error) {
	// 查询抽奖是否存在
	lotteryInfo, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.LotteryId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "lottery_id:%d,err:%v", in.LotteryId, err)
	}
	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("抽奖ID不存在"), "抽奖ID不存在, id: %v", in.LotteryId)
	}

	// 验证抽奖关联的ID是否一致
	if lotteryInfo.ClockTaskId == 0 || lotteryInfo.ClockTaskId != in.ClockTaskId {
		return nil, errors.Wrapf(xerr.NewErrMsg("抽奖关联任务ID不一致"), "抽奖关联任务ID不一致,lotteryInfo.ClockTaskId:%d, in.ClockTaskId:%d", lotteryInfo.ClockTaskId, in.ClockTaskId)
	}

	//todo 验证该用户是否参与抽奖

	// 查询打卡任务是否存在
	clockTaskInfo, err := l.svcCtx.ClockTaskModel.FindOne(l.ctx, in.ClockTaskId)
	if err != nil && !errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "clock_task_id:%d,err:%v", in.ClockTaskId, err)
	}
	if errors.Is(err, model.ErrNotFound) {
		return nil, errors.Wrapf(xerr.NewErrMsg("打卡任务ID不存在"), "打卡任务ID不存在, id:%v", in.ClockTaskId)
	}
	// 判断是否重复完成
	builder := l.svcCtx.ClockTaskRecordModel.SelectBuilder().
		Where(sq.Eq{"user_id": in.UserId}).
		Where(sq.Eq{"lottery_id": lotteryInfo.Id}).
		Where(sq.Eq{"clock_task_id": clockTaskInfo.Id})
	count, err := l.svcCtx.ClockTaskRecordModel.FindCount(l.ctx, builder, "id")
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.Wrapf(xerr.NewErrMsg("用户已完成过打卡任务"), "用户已完成过打卡任务, user_id:%v, lottery_id:%v, clock_task_id:%v", in.UserId, lotteryInfo.Id, clockTaskInfo.Id)
	}

	// 新增完成打卡任务记录
	clockTaskRecord := new(model.ClockTaskRecord)
	clockTaskRecord.LotteryId = lotteryInfo.Id
	clockTaskRecord.UserId = in.UserId
	clockTaskRecord.ClockTaskId = clockTaskInfo.Id

	// 获取中奖倍率
	if clockTaskInfo.ChanceType == constants.Appoint {
		// 指定
		clockTaskRecord.IncreaseMultiple = clockTaskInfo.IncreaseMultiple
	} else {
		// 随机1～10倍
		clockTaskRecord.IncreaseMultiple = rand.Int63n(10) + 1
	}

	insert, err := l.svcCtx.ClockTaskRecordModel.Insert(l.ctx, clockTaskRecord)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTCLOCKTASKRECORD_ERROR), "Lottery Database Exception clockTaskRecord : %+v , err: %v", clockTaskRecord, err)
	}
	clockTaskRecordId, _ := insert.LastInsertId()

	return &pb.AddClockTaskRecordResp{Id: clockTaskRecordId}, nil
}
