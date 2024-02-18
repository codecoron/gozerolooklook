package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"looklook/common/xerr"

	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPrizeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPrizeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPrizeLogic {
	return &SearchPrizeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPrizeLogic) SearchPrize(in *pb.SearchPrizeReq) (*pb.SearchPrizeResp, error) {
	prizes, err := l.svcCtx.PrizeModel.FindPageByLotteryId(l.ctx, in.LotteryId, in.Page, in.Limit)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "Failed to get user's homestay order err : %v , in :%+v", err, in)
	}
	var resp []*pb.Prize
	if len(prizes) > 0 {
		for _, prize := range prizes {
			var pbprize pb.Prize
			_ = copier.Copy(&pbprize, prize)
			resp = append(resp, &pbprize)
		}
	}
	return &pb.SearchPrizeResp{
		Prize: resp,
	}, nil
}
