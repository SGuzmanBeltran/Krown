package tournament

import (
	"context"
	"fmt"
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

func (t *TournamentStore) GetTournaments(req *proto_tournament.GetTournamentsReq) ([]*db.Tournament, error) {
	st := time.Unix(req.StartTime, 0)
	ft := time.Unix(req.FinalTime, 0)

	stTimestamp := &pgtype.Timestamp{
		Time: st.UTC(),
		Valid: true,
	}

	ftTimestamp := &pgtype.Timestamp{
		Time: ft.UTC(),
		Valid: true,
	}

	fmt.Println(stTimestamp.Time.UTC(), ftTimestamp.Time.UTC())
	params := &db.GetTournamentsParams{
		StartTime: *stTimestamp,
		StartTime_2: *ftTimestamp,
	}

	tournaments, err := t.tournamentQueries.GetTournaments(context.Background(), *params)
	if err != nil {
		return nil, err
	}

	var result []*db.Tournament
	for _, dbt := range tournaments {
		result = append(result, &dbt)
	}

	return result, nil
}

func (t *TournamentStore) GetTournament(id int64) (*db.Tournament, error) {
	tournament, err := t.tournamentQueries.GetTournament(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return &tournament, nil
}

func (t *TournamentStore) CreateTournaments(tournaments []db.BatchCreateParams) ([]*db.Tournament, error) {
	batchResults := t.tournamentQueries.BatchCreate(context.Background(), tournaments)
	var tournamentsResults []*db.Tournament
	batchResults.QueryRow(func(i int, t db.Tournament, err error) {
		tournamentsResults = append(tournamentsResults, &t)
	})
	return tournamentsResults, nil;
}