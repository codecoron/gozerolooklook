// Code generated by goctl. DO NOT EDIT.
// Source: checkin.proto

package server

import (
	"context"

	"looklook/app/checkin/cmd/rpc/internal/logic"
	"looklook/app/checkin/cmd/rpc/internal/svc"
	"looklook/app/checkin/cmd/rpc/pb"
)

type CheckinServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedCheckinServer
}

func NewCheckinServer(svcCtx *svc.ServiceContext) *CheckinServer {
	return &CheckinServer{
		svcCtx: svcCtx,
	}
}

// -----------------------checkinRecord-----------------------
func (s *CheckinServer) AddCheckinRecord(ctx context.Context, in *pb.AddCheckinRecordReq) (*pb.AddCheckinRecordResp, error) {
	l := logic.NewAddCheckinRecordLogic(ctx, s.svcCtx)
	return l.AddCheckinRecord(in)
}

func (s *CheckinServer) UpdateCheckinRecord(ctx context.Context, in *pb.UpdateCheckinRecordReq) (*pb.UpdateCheckinRecordResp, error) {
	l := logic.NewUpdateCheckinRecordLogic(ctx, s.svcCtx)
	return l.UpdateCheckinRecord(in)
}

func (s *CheckinServer) GetCheckinRecordByUserId(ctx context.Context, in *pb.GetCheckinRecordByUserIdReq) (*pb.GetCheckinRecordByUserIdResp, error) {
	l := logic.NewGetCheckinRecordByUserIdLogic(ctx, s.svcCtx)
	return l.GetCheckinRecordByUserId(in)
}

// -----------------------integralRecord-----------------------
func (s *CheckinServer) AddIntegral(ctx context.Context, in *pb.AddIntegralReq) (*pb.AddIntegralResp, error) {
	l := logic.NewAddIntegralLogic(ctx, s.svcCtx)
	return l.AddIntegral(in)
}

func (s *CheckinServer) AddIntegralRecord(ctx context.Context, in *pb.AddIntegralRecordReq) (*pb.AddIntegralRecordResp, error) {
	l := logic.NewAddIntegralRecordLogic(ctx, s.svcCtx)
	return l.AddIntegralRecord(in)
}

func (s *CheckinServer) UpdateIntegralRecord(ctx context.Context, in *pb.UpdateIntegralRecordReq) (*pb.UpdateIntegralRecordResp, error) {
	l := logic.NewUpdateIntegralRecordLogic(ctx, s.svcCtx)
	return l.UpdateIntegralRecord(in)
}

func (s *CheckinServer) DelIntegralRecord(ctx context.Context, in *pb.DelIntegralRecordReq) (*pb.DelIntegralRecordResp, error) {
	l := logic.NewDelIntegralRecordLogic(ctx, s.svcCtx)
	return l.DelIntegralRecord(in)
}

func (s *CheckinServer) GetIntegralRecordById(ctx context.Context, in *pb.GetIntegralRecordByIdReq) (*pb.GetIntegralRecordByIdResp, error) {
	l := logic.NewGetIntegralRecordByIdLogic(ctx, s.svcCtx)
	return l.GetIntegralRecordById(in)
}

func (s *CheckinServer) GetIntegralRecordByUserId(ctx context.Context, in *pb.GetIntegralRecordByUserIdReq) (*pb.GetIntegralRecordByUserIdResp, error) {
	l := logic.NewGetIntegralRecordByUserIdLogic(ctx, s.svcCtx)
	return l.GetIntegralRecordByUserId(in)
}

func (s *CheckinServer) SearchIntegralRecord(ctx context.Context, in *pb.SearchIntegralRecordReq) (*pb.SearchIntegralRecordResp, error) {
	l := logic.NewSearchIntegralRecordLogic(ctx, s.svcCtx)
	return l.SearchIntegralRecord(in)
}

// -----------------------taskRecord-----------------------
func (s *CheckinServer) AddTaskRecord(ctx context.Context, in *pb.AddTaskRecordReq) (*pb.AddTaskRecordResp, error) {
	l := logic.NewAddTaskRecordLogic(ctx, s.svcCtx)
	return l.AddTaskRecord(in)
}

func (s *CheckinServer) UpdateTaskRecord(ctx context.Context, in *pb.UpdateTaskRecordReq) (*pb.UpdateTaskRecordResp, error) {
	l := logic.NewUpdateTaskRecordLogic(ctx, s.svcCtx)
	return l.UpdateTaskRecord(in)
}

func (s *CheckinServer) DelTaskRecord(ctx context.Context, in *pb.DelTaskRecordReq) (*pb.DelTaskRecordResp, error) {
	l := logic.NewDelTaskRecordLogic(ctx, s.svcCtx)
	return l.DelTaskRecord(in)
}

func (s *CheckinServer) GetTaskRecordById(ctx context.Context, in *pb.GetTaskRecordByIdReq) (*pb.GetTaskRecordByIdResp, error) {
	l := logic.NewGetTaskRecordByIdLogic(ctx, s.svcCtx)
	return l.GetTaskRecordById(in)
}

func (s *CheckinServer) GetTaskRecordByUserId(ctx context.Context, in *pb.GetTaskRecordByUserIdReq) (*pb.GetTaskRecordByUserIdResp, error) {
	l := logic.NewGetTaskRecordByUserIdLogic(ctx, s.svcCtx)
	return l.GetTaskRecordByUserId(in)
}

func (s *CheckinServer) SearchTaskRecord(ctx context.Context, in *pb.SearchTaskRecordReq) (*pb.SearchTaskRecordResp, error) {
	l := logic.NewSearchTaskRecordLogic(ctx, s.svcCtx)
	return l.SearchTaskRecord(in)
}

func (s *CheckinServer) GetTaskProgress(ctx context.Context, in *pb.GetTaskProgressReq) (*pb.GetTaskProgressResp, error) {
	l := logic.NewGetTaskProgressLogic(ctx, s.svcCtx)
	return l.GetTaskProgress(in)
}

// -----------------------tasks-----------------------
func (s *CheckinServer) AddTasks(ctx context.Context, in *pb.AddTasksReq) (*pb.AddTasksResp, error) {
	l := logic.NewAddTasksLogic(ctx, s.svcCtx)
	return l.AddTasks(in)
}

func (s *CheckinServer) UpdateTasks(ctx context.Context, in *pb.UpdateTasksReq) (*pb.UpdateTasksResp, error) {
	l := logic.NewUpdateTasksLogic(ctx, s.svcCtx)
	return l.UpdateTasks(in)
}

func (s *CheckinServer) DelTasks(ctx context.Context, in *pb.DelTasksReq) (*pb.DelTasksResp, error) {
	l := logic.NewDelTasksLogic(ctx, s.svcCtx)
	return l.DelTasks(in)
}

func (s *CheckinServer) GetTasksById(ctx context.Context, in *pb.GetTasksByIdReq) (*pb.GetTasksByIdResp, error) {
	l := logic.NewGetTasksByIdLogic(ctx, s.svcCtx)
	return l.GetTasksById(in)
}

func (s *CheckinServer) SearchTasks(ctx context.Context, in *pb.SearchTasksReq) (*pb.SearchTasksResp, error) {
	l := logic.NewSearchTasksLogic(ctx, s.svcCtx)
	return l.SearchTasks(in)
}
