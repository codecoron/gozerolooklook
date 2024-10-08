// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	lottery "looklook/app/lottery/cmd/api/internal/handler/lottery"
	"looklook/app/lottery/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/lottery/TestValidator",
				Handler: lottery.TestValidatorHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/lotteryList",
				Handler: lottery.LotteryListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/participations",
				Handler: lottery.SearchParticipationHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/lottery/clockTaskTypeList",
				Handler: lottery.ClockTaskTypeListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/lottery/chanceTypeList",
				Handler: lottery.ChanceTypeListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/getLotteryWinnersList",
				Handler: lottery.GetLotteryWinList2Handler(serverCtx),
			},
		},
		rest.WithPrefix("/lottery/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/lottery/createLottery",
				Handler: lottery.CreateLotteryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/updateLottery",
				Handler: lottery.UpdateLotteryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/getLotteryWinList",
				Handler: lottery.GetLotteryWinListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/lotteryDetail",
				Handler: lottery.LotteryDetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/CheckIsParticipated",
				Handler: lottery.CheckIsParticipatedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/participation",
				Handler: lottery.AddLotteryParticipationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/setLotteryIsSelected",
				Handler: lottery.SetLotteryIsSelectedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/lotteryListAfterLogin",
				Handler: lottery.LotteryListAfterLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/checkIsWin",
				Handler: lottery.CheckIsWinHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/createClockTaskRecord",
				Handler: lottery.CreateClockTaskRecordHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/lottery/getLotteryListByUserId",
				Handler: lottery.GetLotteryListByUserIdHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.JwtAuth.AccessSecret),
		rest.WithPrefix("/lottery/v1"),
	)
}
