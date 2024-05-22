package types

import (
	"context"
	proto_tournament "krown/services/genproto/tournament"
)

type TournamentService interface {
	GetTournaments(context.Context, *proto_tournament.GetTournamentsReq) (*proto_tournament.GetTournamentsRes, error)
	GetTournament(context.Context, *proto_tournament.GetTournamentReq) (*proto_tournament.GetTournamentRes, error)
}
