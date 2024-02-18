package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/app/usercenter/cmd/rpc/usercenter"
	"looklook/common/xerr"
)

type SetIsSelectedLotteryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetIsSelectedLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetIsSelectedLotteryLogic {
	return &SetIsSelectedLotteryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetIsSelectedLotteryLogic) SetIsSelectedLottery(in *pb.SetIsSelectedLotteryReq) (*pb.SetIsSelectedLotteryResp, error) {
	var IsSelected int64
	//添加事务处理
	err := l.svcCtx.LotteryModel.Trans(l.ctx, func(context context.Context, session sqlx.Session) error {
		// 1. 找到user，确认是否为admin;如果没找到，说明用户不存在。调用userCenter下的操作。
		userinfo, err := l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.GetUserInfoReq{
			Id: in.UserId,
		})
		if err != nil {
			return err
		}
		// 2. 如果是，则找到id对应的抽奖，将is_selected字段取反
		if userinfo.User.IsAdmin != 0 {
			lottery, err := l.svcCtx.LotteryModel.FindOne(l.ctx, in.Id)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_GET_LOTTERY_BYLOTTERYID_ERROR), "SetIsSelectedLottery, id:%v, error: %v", in.Id, err)
			}
			if lottery.IsSelected == 1 {
				lottery.IsSelected = 0
			} else {
				lottery.IsSelected = 1
				IsSelected = 1
			}
			err = l.svcCtx.LotteryModel.Update(l.ctx, lottery)
			if err != nil {
				return errors.Wrapf(xerr.NewErrCode(xerr.DB_UPDATE_LOTTERY_ERROR), "SetIsSelectedLottery, lottery:%v, error: %v", lottery, err)
			}
		} else {
			return errors.Wrapf(xerr.NewErrCode(xerr.DB_NO_SET_LOTTERY_ISSELECT_PERMISSION_ERROR), "SetIsSelectedLottery, userinfo.User.IsAdmin:%v, error: %v", userinfo.User.IsAdmin, err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &pb.SetIsSelectedLotteryResp{
		IsSelected: IsSelected,
	}, nil
}
