package tournament

import (
	"context"
	"fmt"
	"krown/db"
	proto_tournament "krown/services/genproto/tournament"
)

type TournamentService struct {
	tournamentStore *TournamentStore
}

func NewTournamentService(tournamentStore *TournamentStore) *TournamentService {
	return &TournamentService{tournamentStore}
}

func (t *TournamentService) GetTournaments(c context.Context, req *proto_tournament.GetTournamentsReq) (*proto_tournament.GetTournamentsRes, error) {
	dbTournaments, err := t.tournamentStore.GetTournaments(req)
	if err != nil {
		fmt.Println("Error")
	}
	protoTournatments, err := convertDBTournamentsToProtoTournaments(dbTournaments)
	if err != nil {
		fmt.Println("Error, converting tournaments")
	}
	tournaments := &proto_tournament.GetTournamentsRes{
		Tournaments: protoTournatments,
	}
	return tournaments, nil
}

func (t *TournamentService) GetTournament(c context.Context, req *proto_tournament.GetTournamentReq) (*proto_tournament.GetTournamentRes, error) {
	dbTournament, err := t.tournamentStore.GetTournament(req.TournamentId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	protoTournament, err := convertDBTournamentToProto(*dbTournament)
	if err != nil {
		return nil, err
	}
	tournaments := &proto_tournament.GetTournamentRes{
		Tournament: protoTournament,
	}
	return tournaments, nil
}

func convertDBTournamentsToProtoTournaments(dbtournaments []db.Tournament) ([]*proto_tournament.Tournament, error) {
	var protoTournaments []*proto_tournament.Tournament

	for _, dbt := range dbtournaments {
		// Convert pgtype.Timestamp to int64
		pt, err := convertDBTournamentToProto(dbt)
		if err != nil {
			return nil, err
		}
		protoTournaments = append(protoTournaments, pt)
	}

	return protoTournaments, nil
}

func convertDBTournamentToProto(dbtournament db.Tournament) (*proto_tournament.Tournament, error) {
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
