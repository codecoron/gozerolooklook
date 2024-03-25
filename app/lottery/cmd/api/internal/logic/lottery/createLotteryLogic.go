package lottery

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/app/lottery/cmd/api/internal/svc"
	"looklook/app/lottery/cmd/api/internal/types"
	"looklook/app/lottery/cmd/rpc/lottery"
	"looklook/app/lottery/cmd/rpc/pb"
	"looklook/common/ctxdata"
	"looklook/common/xerr"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLotteryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLotteryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLotteryLogic {
	return &CreateLotteryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLotteryLogic) CreateLottery(req *types.CreateLotteryReq) (resp *types.CreateLotteryResp, err error) {
	userId := ctxdata.GetUidFromCtx(l.ctx)
	//循环赋值 TODO 寻找更好的方案
	var pbPrizes []*pb.Prize
	for _, reqPrize := range req.Prizes {
		pbPrize := new(pb.Prize)
		err := copier.Copy(&pbPrize, reqPrize)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("Copy pbPrize Error"), "Copy pbPrize Error req: %+v , err : %v ", pbPrize, err)
		}
		pbPrizes = append(pbPrizes, pbPrize)
	}
	pbClockTask := new(pb.ClockTask)
	if req.IsClocked == 1 && req.ClockTask != nil {
		err = copier.Copy(&pbClockTask, req.ClockTask)
		if err != nil {
			return nil, errors.Wrapf(xerr.NewErrMsg("Copy pbClockTask Error"), "Copy pbClockTask Error req: %+v , err : %v ", pbClockTask, err)
		}
	} else {
		pbClockTask = nil
	}

	addLottery, err := l.svcCtx.LotteryRpc.AddLottery(l.ctx, &lottery.AddLotteryReq{
		UserId:        userId,
		Name:          req.Name,
		Thumb:         req.Thumb,
		AnnounceType:  req.AnnounceType,
		AnnounceTime:  req.AnnounceTime,
		JoinNumber:    req.JoinNumber,
		Introduce:     req.Introduce,
		AwardDeadline: req.AwardDeadline,
		Prizes:        pbPrizes,
		SponsorId:     req.SponsorId,
		IsClocked:     req.IsClocked,
		ClockTask:     pbClockTask,
		PublishType:   req.PublishType,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateLotteryResp{
		Id: addLottery.Id,
	}, nil
}

// 这里用到了什么设计模式
// 这里用到了抽象工厂模式
// 代码中的抽象工厂模式的实现：
// 1. CreateLotteryLogic 是抽象工厂
// 2. NewCreateLotteryLogic 是具体工厂
// 3. svc.ServiceContext 是抽象产品
// 4. LotteryRpc 是具体产品
// 5. 抽象工厂模式的实现中，抽象工厂和具体工厂都实现了同一个接口 CreateLotteryLogic
// 这里用到了代理模式
// 代码中的代理模式：
// 1. CreateLotteryLogic是代理主题角色
// 2. svc.ServiceContext是真实主题角色
// 3. CreateLotteryLogic实现了CreateLottery方法，这个方法的实现是调用了svc.ServiceContext的LotteryRpc.AddLottery方法
// 4. CreateLotteryLogic和svc.ServiceContext都实现了同一个接口
// 5. CreateLotteryLogic保存了一个引用使得代理可以访问实体
// 6. CreateLotteryLogic提供了一个与svc.ServiceContext相同的接口，这样代理就可以替代实体
// 7. CreateLotteryLogic控制对svc.ServiceContext的存取
// 8. CreateLotteryLogic提供了其他功能，比如记录日志

// 哪里可能会用到单例模式
// 1. 服务端的配置信息
// 2. 数据库连接池
// 3. 日志对象
// 4. 线程池
// 5. 缓存
// 6. 对象池
// 7. 任务管理器
// 8. 连接池
// 9. 等等
// 代码中的单例模式的实现：
// 1. svc.ServiceContext中的Config是单例模式
