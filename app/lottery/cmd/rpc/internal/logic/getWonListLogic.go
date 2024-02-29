package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWonListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWonListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWonListLogic {
	return &GetWonListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetWonListLogic) GetWonList(in *pb.GetWonListReq) (*pb.GetWonListResp, error) {
	res, err := l.svcCtx.LotteryParticipationModel.GetWonListByUserId(l.ctx, in.UserId, in.Page, in.Size, in.LastId)
	if err != nil {
		return nil, err
	}
	var list []*pb.WonList
	for _, item := range res {
		pbWonList := new(pb.WonList)
		pbWonList.Id = item.Id
		pbWonList.LotteryId = item.LotteryId
		pbWonList.UserId = item.UserId
		pbWonList.IsWon = true
		prize, err := l.svcCtx.PrizeModel.FindOne(l.ctx, item.PrizeId)
		if err != nil {
			return nil, err
		}
		pbWonList.Prize = new(pb.Prize)
		err = copier.Copy(pbWonList.Prize, prize)
		if err != nil {
			return nil, err
		}
		list = append(list, pbWonList)
	}

	return &pb.GetWonListResp{
		List: list,
	}, nil
}
