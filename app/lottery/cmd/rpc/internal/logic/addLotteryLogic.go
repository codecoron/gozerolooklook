package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/lottery/model"
	"looklook/common/xerr"
	"time"
)

type AddLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLotteryLogic {
	return &AddLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------发起抽奖----------------------
func (l *AddLotteryLogic) AddLottery(in *pb.AddLotteryReq) (*pb.AddLotteryResp, error) {
	var lotteryId int64
	//添加事务处理
	err := l.svcCtx.LotteryModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		//抽奖基本信息
		lottery := new(model.Lottery)
		lottery.UserId = in.UserId
		lottery.Name = in.Name
		lottery.AwardDeadline = time.Unix(in.AwardDeadline, 0)
		lottery.Introduce = in.Introduce
		lottery.JoinNumber = in.JoinNumber
		lottery.AnnounceType = in.AnnounceType
		lottery.AnnounceTime = time.Unix(in.AnnounceTime, 0)
		lottery.Thumb = in.Thumb
		lottery.IsSelected = 0
		lottery.IsAnnounced = 0
		//打印出sql 调试错误
		insert, err := l.svcCtx.LotteryModel.TransInsert(l.ctx, session, lottery)
		if err != nil {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTLOTTERY_ERROR), "Lottery Database Exception lottery : %+v , err: %v", lottery, err)
		}
		lotteryId, _ = insert.LastInsertId()
		//添加奖品信息
		for _, pbPrize := range in.Prizes {
			prize := new(model.Prize)
			err := copier.Copy(&prize, pbPrize)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTPRIZE_ERROR), "copier : %+v , err: %v", prize, err)
			}
			prize.LotteryId = lotteryId
			_, err = l.svcCtx.PrizeModel.TransInsert(l.ctx, session, prize)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_INSERTPRIZE_ERROR), "Lottery Database Exception prize : %+v , err: %v", prize, err)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &pb.AddLotteryResp{
		Id: lotteryId,
	}, nil
}
