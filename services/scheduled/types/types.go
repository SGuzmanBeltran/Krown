package types

import (
	"context"
	"krown/db"
	proto_scheduled "krown/services/genproto/scheduled"
	proto_tournament "krown/services/genproto/tournament"
)

type ScheduledService interface {
	GetScheduleds(context.Context, *proto_scheduled.GetScheduledsReq) (*proto_scheduled.GetScheduledsRes, error)
	GetScheduled(context.Context, *proto_scheduled.GetScheduledReq) (*proto_scheduled.GetScheduledRes, error)
	CreateScheduled(context.Context, *proto_scheduled.CreateScheduledReq) (*proto_scheduled.CreateScheduledRes, error)
}

func ConvertDBCreateScheduledToProtoTournaments(dbtournaments []db.ScheduledTournament) []*proto_tournament.CreateTournamentDTO {
	var protoTournaments []*proto_tournament.CreateTournamentDTO
	for _, dbt := range dbtournaments {
		pt := ConvertDBCreateScheduledToProto(&dbt)
		protoTournaments = append(protoTournaments, pt)
	}
	return protoTournaments
}

func ConvertDBCreateScheduledToProto(dbtournament *db.ScheduledTournament) *proto_tournament.CreateTournamentDTO {
	var protoTournaments *proto_tournament.CreateTournamentDTO
	startTime := dbtournament.StartTime.Time.Unix()
	protoTournaments = &proto_tournament.CreateTournamentDTO{
		Name:      dbtournament.Name,
		EntryFee:  dbtournament.EntryFee,
		StartTime: startTime,
	}
	return protoTournaments
}