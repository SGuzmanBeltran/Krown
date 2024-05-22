package types

import (
	"context"
	"krown/db"
	proto_tournament "krown/services/genproto/tournament"
)

type TournamentService interface {
	GetTournaments(context.Context, *proto_tournament.GetTournamentsReq) (*proto_tournament.GetTournamentsRes, error)
	GetTournament(context.Context, *proto_tournament.GetTournamentReq) (*proto_tournament.GetTournamentRes, error)
}

func ConvertDBTournamentsToProtoTournaments(dbtournaments []db.Tournament) ([]*proto_tournament.Tournament, error) {
	var protoTournaments []*proto_tournament.Tournament

	for _, dbt := range dbtournaments {
		// Convert pgtype.Timestamp to int64
		pt, err := ConvertDBTournamentToProto(dbt)
		if err != nil {
			return nil, err
		}
		protoTournaments = append(protoTournaments, pt)
	}

	return protoTournaments, nil
}

func ConvertDBTournamentToProto(dbtournament db.Tournament) (*proto_tournament.Tournament, error) {
	var protoTournaments *proto_tournament.Tournament

	// Convert pgtype.Timestamp to int64
	startTime := dbtournament.StartTime.Time.Unix()

	protoTournaments = &proto_tournament.Tournament{
		Id:        dbtournament.ID,
		Name:      dbtournament.Name,
		EntryFee:  int64(dbtournament.EntryFee),
		StartTime: startTime, // Ensure startTime is in milliseconds
	}

	return protoTournaments, nil
}
