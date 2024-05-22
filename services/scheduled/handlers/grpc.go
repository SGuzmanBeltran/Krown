package handlers

import (
	"context"
	proto_scheduled "krown/services/genproto/scheduled"
	"krown/services/scheduled/types"

	"google.golang.org/grpc"
)

type ScheduledGrpcHandler struct {
	scheduledService types.ScheduledService
	proto_scheduled.UnimplementedScheduledServiceServer
}

func NewGrpcScheduledService(grpc *grpc.Server, scheduledService types.ScheduledService) {
	gRpcHadler := &ScheduledGrpcHandler{scheduledService: scheduledService}
	proto_scheduled.RegisterScheduledServiceServer(grpc, gRpcHadler)
}

func (h *ScheduledGrpcHandler) GetScheduleds(ctx context.Context, req *proto_scheduled.GetScheduledsReq) (*proto_scheduled.GetScheduledsRes, error) {
	scheduleds, err := h.scheduledService.GetScheduleds(ctx, req)
	if (err != nil) {
		return nil, err
	}
	return scheduleds, nil
}

func (h *ScheduledGrpcHandler) GetScheduled(ctx context.Context, req *proto_scheduled.GetScheduledReq) (*proto_scheduled.GetScheduledRes, error) {
	scheduled, err := h.scheduledService.GetScheduled(ctx, req)
	if (err != nil) {
		return nil, err
	}
	return scheduled, nil
}

func (h *ScheduledGrpcHandler) CreateScheduled(ctx context.Context, req *proto_scheduled.CreateScheduledReq) (*proto_scheduled.CreateScheduledRes, error) {
	scheduled, err := h.scheduledService.CreateScheduled(ctx, req)
	if (err != nil) {
		return nil, err
	}
	return scheduled, nil
}