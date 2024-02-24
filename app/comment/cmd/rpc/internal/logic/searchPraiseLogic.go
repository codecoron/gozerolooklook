package logic

import (
	"context"

	"looklook/app/comment/cmd/rpc/internal/svc"
	"looklook/app/comment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPraiseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPraiseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPraiseLogic {
	return &SearchPraiseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPraiseLogic) SearchPraise(in *pb.SearchPraiseReq) (*pb.SearchPraiseResp, error) {
	list, err := l.svcCtx.PraiseModel.PraiseList(l.ctx, in.Page, in.Limit, in.LastId)
	if err != nil {
		return nil, err
	}
	var resp []*pb.Praise
	if len(list) > 0 {
		for _, praise := range list {
			var pbPraise pb.Praise
			pbPraise.Id = praise.Id
			pbPraise.UserId = praise.UserId
			pbPraise.CommentId = praise.CommentId
			resp = append(resp, &pbPraise)
		}
	}

	return &pb.SearchPraiseResp{
		Praise: resp,
	}, nil
}
