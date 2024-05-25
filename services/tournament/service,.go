package tournament

import (
	"context"
	"fmt"
	"krown/db"
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
	protoTournament, err := types.ConvertDBTournamentToProto(dbTournament)
	if err != nil {
		return nil, err
	}
	tournaments := &proto_tournament.GetTournamentRes{
		Tournament: protoTournament,
	}
	return tournaments, nil
}

func (t *TournamentService) CreateTournaments(c context.Context, tournaments []*types.CreateTournament) (*proto_tournament.CreateTournamentRes, error) {
	var storeTournamentParams []db.BatchCreateParams
	for _, dbt := range tournaments {
		pt := db.BatchCreateParams(*dbt)
		storeTournamentParams = append(storeTournamentParams, pt)
	}
	createTournaments, err := t.tournamentStore.CreateTournaments(storeTournamentParams)
	if err != nil {
		return nil, err
	}
	protoTournaments, err := types.ConvertDBTournamentsToProtoTournaments(createTournaments)
	if err != nil {
		return nil, err
	}
	response := &proto_tournament.CreateTournamentRes{
		Tournaments: protoTournaments,
	}
	return response, nil
}
