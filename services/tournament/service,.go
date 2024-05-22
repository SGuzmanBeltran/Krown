package tournament

import (
	"context"
	"fmt"
	proto_tournament "krown/services/genproto/tournament"
	"krown/services/tournament/types"
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
	protoTournatments, err := types.ConvertDBTournamentsToProtoTournaments(dbTournaments)
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
	protoTournament, err := types.ConvertDBTournamentToProto(*dbTournament)
	if err != nil {
		return nil, err
	}
	tournaments := &proto_tournament.GetTournamentRes{
		Tournament: protoTournament,
	}
	return tournaments, nil
}
