package scheduled

import (
	"context"
	proto_scheduled "krown/services/genproto/scheduled"
	"krown/services/scheduled/types"
)

type ScheduledService struct {
	scheduledStore *ScheduledStore
}

func NewScheduledService(scheduledStore *ScheduledStore) *ScheduledService {
	return &ScheduledService{scheduledStore}
}

func (s *ScheduledService) GetScheduleds(c context.Context, req *proto_scheduled.GetScheduledsReq) (*proto_scheduled.GetScheduledsRes, error) {
	dbScheduled, err := s.scheduledStore.GetScheduleds()
	if err != nil {
		return nil, err
	}

	protoScheduled, err := types.ConvertDBScheduledToProtos(dbScheduled)
	if err != nil {
		return nil, err
	}

	scheduled := &proto_scheduled.GetScheduledsRes{
		Scheduleds: protoScheduled,
	}
	return scheduled, nil
}

func (s *ScheduledService) GetScheduled(c context.Context, req *proto_scheduled.GetScheduledReq) (*proto_scheduled.GetScheduledRes, error) {
	dbScheduled, err := s.scheduledStore.GetScheduled(req.ScheduledId)
	if err != nil {
		return nil, err
	}

	protoScheduled, err := types.ConvertDBScheduledToProto(*dbScheduled)
	if err != nil {
		return nil, err
	}

	scheduled := &proto_scheduled.GetScheduledRes{
		Scheduled: protoScheduled,
	}
	return scheduled, nil
}

func (s *ScheduledService) CreateScheduled(c context.Context, req *proto_scheduled.CreateScheduledReq) (*proto_scheduled.CreateScheduledRes, error) {
	scheduledParams := types.ConvertProtoToParams(req.Scheduled)
	createdScheduled, err := s.scheduledStore.CreateScheduled(scheduledParams)
	if err != nil {
		return nil, err
	}
	scheduledProto, err := types.ConvertDBScheduledToProto(*createdScheduled)
	if err != nil {
		return nil, err
	}
	scheduled := &proto_scheduled.CreateScheduledRes{
		Scheduled: scheduledProto,
	}
	return scheduled, nil
}
