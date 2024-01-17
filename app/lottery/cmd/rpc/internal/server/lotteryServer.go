// Code generated by goctl. DO NOT EDIT.
// Source: lottery.proto

package server

import (
	"context"

	"looklook/app/lottery/cmd/rpc/internal/logic"
	"looklook/app/lottery/cmd/rpc/internal/svc"
	"looklook/app/lottery/cmd/rpc/pb"
)

type LotteryServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedLotteryServer
}

func NewLotteryServer(svcCtx *svc.ServiceContext) *LotteryServer {
	return &LotteryServer{
		svcCtx: svcCtx,
	}
}

// -----------------------抽奖表-----------------------
func (s *LotteryServer) AddLottery(ctx context.Context, in *pb.AddLotteryReq) (*pb.AddLotteryResp, error) {
	l := logic.NewAddLotteryLogic(ctx, s.svcCtx)
	return l.AddLottery(in)
}

func (s *LotteryServer) UpdateLottery(ctx context.Context, in *pb.UpdateLotteryReq) (*pb.UpdateLotteryResp, error) {
	l := logic.NewUpdateLotteryLogic(ctx, s.svcCtx)
	return l.UpdateLottery(in)
}

func (s *LotteryServer) DelLottery(ctx context.Context, in *pb.DelLotteryReq) (*pb.DelLotteryResp, error) {
	l := logic.NewDelLotteryLogic(ctx, s.svcCtx)
	return l.DelLottery(in)
}

func (s *LotteryServer) GetLotteryById(ctx context.Context, in *pb.GetLotteryByIdReq) (*pb.GetLotteryByIdResp, error) {
	l := logic.NewGetLotteryByIdLogic(ctx, s.svcCtx)
	return l.GetLotteryById(in)
}

func (s *LotteryServer) SearchLottery(ctx context.Context, in *pb.SearchLotteryReq) (*pb.SearchLotteryResp, error) {
	l := logic.NewSearchLotteryLogic(ctx, s.svcCtx)
	return l.SearchLottery(in)
}

func (s *LotteryServer) SearchIsSelectedLottery(ctx context.Context, in *pb.SearchIsSelectedLotteryReq) (*pb.SearchIsSelectedLotteryResp, error) {
	l := logic.NewSearchIsSelectedLotteryLogic(ctx, s.svcCtx)
	return l.SearchIsSelectedLottery(in)
}

// -----------------------奖品表-----------------------
func (s *LotteryServer) AddPrize(ctx context.Context, in *pb.AddPrizeReq) (*pb.AddPrizeResp, error) {
	l := logic.NewAddPrizeLogic(ctx, s.svcCtx)
	return l.AddPrize(in)
}

func (s *LotteryServer) UpdatePrize(ctx context.Context, in *pb.UpdatePrizeReq) (*pb.UpdatePrizeResp, error) {
	l := logic.NewUpdatePrizeLogic(ctx, s.svcCtx)
	return l.UpdatePrize(in)
}

func (s *LotteryServer) DelPrize(ctx context.Context, in *pb.DelPrizeReq) (*pb.DelPrizeResp, error) {
	l := logic.NewDelPrizeLogic(ctx, s.svcCtx)
	return l.DelPrize(in)
}

func (s *LotteryServer) GetPrizeById(ctx context.Context, in *pb.GetPrizeByIdReq) (*pb.GetPrizeByIdResp, error) {
	l := logic.NewGetPrizeByIdLogic(ctx, s.svcCtx)
	return l.GetPrizeById(in)
}

func (s *LotteryServer) SearchPrize(ctx context.Context, in *pb.SearchPrizeReq) (*pb.SearchPrizeResp, error) {
	l := logic.NewSearchPrizeLogic(ctx, s.svcCtx)
	return l.SearchPrize(in)
}
