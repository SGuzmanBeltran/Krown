package tournament

import (
	"context"
	"krown/db"
	proto_tournament "krown/services/genproto/tournament"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type TournamentStore struct {
	tournamentQueries *db.Queries
}

func NewTournamentStore(tournamentQueries *db.Queries) *TournamentStore {
	return &TournamentStore{tournamentQueries}
}

func (t *TournamentStore) GetTournaments(req *proto_tournament.GetTournamentsReq) ([]db.Tournament, error) {
	st := time.Unix(req.StartTime, 0)
	ft := time.Unix(req.FinalTime, 0)

	stTimestamp := &pgtype.Timestamp{
		Time: st,
		Valid: true,
	}

	ftTimestamp := &pgtype.Timestamp{
		Time: ft,
		Valid: true,
	}

	params := &db.GetTournamentsParams{
		StartTime: *stTimestamp,
		StartTime_2: *ftTimestamp,
	}

	tournaments, err := t.tournamentQueries.GetTournaments(context.Background(), *params)
	if err != nil {
		return nil, err
	}

	return tournaments, nil
}

func (t *TournamentStore) GetTournament(id int64) (*db.Tournament, error) {
	tournament, err := t.tournamentQueries.GetTournament(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &tournament, nil
}