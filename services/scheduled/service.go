package scheduled

import (
	"context"
	"fmt"
	proto_scheduled "krown/services/genproto/scheduled"
	proto_tournament "krown/services/genproto/tournament"
	"krown/services/scheduled/types"

	"google.golang.org/grpc"
)

type ScheduledService struct {
	scheduledStore *ScheduledStore
	tournamentConn *grpc.ClientConn
}

func NewScheduledService(scheduledStore *ScheduledStore, tournamentConn *grpc.ClientConn) *ScheduledService {
	return &ScheduledService{scheduledStore, tournamentConn}
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

func (s *ScheduledService) GenerateTournaments(date int64) {
	createdScheduled, err := s.scheduledStore.GetScheduledsByStartTime(date)
	if err != nil {
		fmt.Println("Error getting scheduled tournaments")
	}
	tClient := proto_tournament.NewTournamentServiceClient(s.tournamentConn)
	protoCreateTournaments := types.ConvertDBCreateScheduledToProtoTournaments(createdScheduled)
	request := &proto_tournament.CreateTournamentReq{
		CreateTournaments: protoCreateTournaments,
	}
	tClient.CreateTournaments(context.Background(), request)
}
