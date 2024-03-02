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

func (s *LotteryServer) SetIsSelectedLottery(ctx context.Context, in *pb.SetIsSelectedLotteryReq) (*pb.SetIsSelectedLotteryResp, error) {
	l := logic.NewSetIsSelectedLotteryLogic(ctx, s.svcCtx)
	return l.SetIsSelectedLottery(in)
}

func (s *LotteryServer) LotteryDetail(ctx context.Context, in *pb.LotteryDetailReq) (*pb.LotteryDetailResp, error) {
	l := logic.NewLotteryDetailLogic(ctx, s.svcCtx)
	return l.LotteryDetail(in)
}

func (s *LotteryServer) LotterySponsor(ctx context.Context, in *pb.LotterySponsorReq) (*pb.LotterySponsorResp, error) {
	l := logic.NewLotterySponsorLogic(ctx, s.svcCtx)
	return l.LotterySponsor(in)
}

func (s *LotteryServer) AnnounceLottery(ctx context.Context, in *pb.AnnounceLotteryReq) (*pb.AnnounceLotteryResp, error) {
	l := logic.NewAnnounceLotteryLogic(ctx, s.svcCtx)
	return l.AnnounceLottery(in)
}

func (s *LotteryServer) CheckUserCreatedLottery(ctx context.Context, in *pb.CheckUserCreatedLotteryReq) (*pb.CheckUserCreatedLotteryResp, error) {
	l := logic.NewCheckUserCreatedLotteryLogic(ctx, s.svcCtx)
	return l.CheckUserCreatedLottery(in)
}

func (s *LotteryServer) CheckUserCreatedLotteryAndToday(ctx context.Context, in *pb.CheckUserCreatedLotteryAndTodayReq) (*pb.CheckUserCreatedLotteryAndTodayResp, error) {
	l := logic.NewCheckUserCreatedLotteryAndTodayLogic(ctx, s.svcCtx)
	return l.CheckUserCreatedLotteryAndToday(in)
}

func (s *LotteryServer) CheckUserCreatedLotteryAndThisWeek(ctx context.Context, in *pb.CheckUserCreatedLotteryAndThisWeekReq) (*pb.CheckUserCreatedLotteryAndThisWeekResp, error) {
	l := logic.NewCheckUserCreatedLotteryAndThisWeekLogic(ctx, s.svcCtx)
	return l.CheckUserCreatedLotteryAndThisWeek(in)
}

func (s *LotteryServer) GetLotteryListAfterLogin(ctx context.Context, in *pb.GetLotteryListAfterLoginReq) (*pb.GetLotteryListAfterLoginResp, error) {
	l := logic.NewGetLotteryListAfterLoginLogic(ctx, s.svcCtx)
	return l.GetLotteryListAfterLogin(in)
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

func (s *LotteryServer) GetPrizeListByLotteryId(ctx context.Context, in *pb.GetPrizeListByLotteryIdReq) (*pb.GetPrizeListByLotteryIdResp, error) {
	l := logic.NewGetPrizeListByLotteryIdLogic(ctx, s.svcCtx)
	return l.GetPrizeListByLotteryId(in)
}

// -----------------------参与抽奖-----------------------
func (s *LotteryServer) AddLotteryParticipation(ctx context.Context, in *pb.AddLotteryParticipationReq) (*pb.AddLotteryParticipationResp, error) {
	l := logic.NewAddLotteryParticipationLogic(ctx, s.svcCtx)
	return l.AddLotteryParticipation(in)
}

func (s *LotteryServer) SearchLotteryParticipation(ctx context.Context, in *pb.SearchLotteryParticipationReq) (*pb.SearchLotteryParticipationResp, error) {
	l := logic.NewSearchLotteryParticipationLogic(ctx, s.svcCtx)
	return l.SearchLotteryParticipation(in)
}

func (s *LotteryServer) GetParticipationUserIdsByLotteryId(ctx context.Context, in *pb.GetParticipationUserIdsByLotteryIdReq) (*pb.GetParticipationUserIdsByLotteryIdResp, error) {
	l := logic.NewGetParticipationUserIdsByLotteryIdLogic(ctx, s.svcCtx)
	return l.GetParticipationUserIdsByLotteryId(in)
}

func (s *LotteryServer) CheckIsParticipated(ctx context.Context, in *pb.CheckIsParticipatedReq) (*pb.CheckIsParticipatedResp, error) {
	l := logic.NewCheckIsParticipatedLogic(ctx, s.svcCtx)
	return l.CheckIsParticipated(in)
}

func (s *LotteryServer) GetSelectedLotteryStatistic(ctx context.Context, in *pb.GetSelectedLotteryStatisticReq) (*pb.GetSelectedLotteryStatisticResp, error) {
	l := logic.NewGetSelectedLotteryStatisticLogic(ctx, s.svcCtx)
	return l.GetSelectedLotteryStatistic(in)
}

func (s *LotteryServer) CheckSelectedLotteryParticipated(ctx context.Context, in *pb.CheckSelectedLotteryParticipatedReq) (*pb.CheckSelectedLotteryParticipatedResp, error) {
	l := logic.NewCheckSelectedLotteryParticipatedLogic(ctx, s.svcCtx)
	return l.CheckSelectedLotteryParticipated(in)
}

func (s *LotteryServer) CheckUserIsWon(ctx context.Context, in *pb.CheckUserIsWonReq) (*pb.CheckUserIsWonResp, error) {
	l := logic.NewCheckUserIsWonLogic(ctx, s.svcCtx)
	return l.CheckUserIsWon(in)
}

func (s *LotteryServer) GetWonList(ctx context.Context, in *pb.GetWonListReq) (*pb.GetWonListResp, error) {
	l := logic.NewGetWonListLogic(ctx, s.svcCtx)
	return l.GetWonList(in)
}

func (s *LotteryServer) GetWonListCount(ctx context.Context, in *pb.GetWonListCountReq) (*pb.GetWonListCountResp, error) {
	l := logic.NewGetWonListCountLogic(ctx, s.svcCtx)
	return l.GetWonListCount(in)
}
