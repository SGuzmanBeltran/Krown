package scheduled

import (
	"context"
	"krown/db"
)

type ScheduledStore struct {
	queries *db.Queries
}

func NewScheduledStore(queries *db.Queries) *ScheduledStore {
	return &ScheduledStore{queries: queries}
}

func (s *ScheduledStore) GetScheduleds() ([]db.ScheduledTournament, error) {
	scheduled, err := s.queries.GetScheduledTournaments(context.Background())
	if err != nil {
		return nil, err
	}

	return scheduled, nil
}

func (s *ScheduledStore) GetScheduled(id int64) (*db.ScheduledTournament, error) {
	scheduled, err := s.queries.GetScheduledTournament(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &scheduled, nil
}

func (s *ScheduledStore) CreateScheduled(scheduled *db.CreateScheduledTournamentParams) (*db.ScheduledTournament, error) {
	newScheduled, err := s.queries.CreateScheduledTournament(context.Background(), *scheduled)
	if err != nil {
		return nil, err
	}
	return &newScheduled, nil
}

func (s *ScheduledStore) GetScheduledsByStartTime(date int64) ([]db.ScheduledTournament, error) {
	scheduled, err := s.queries.GetScheduledTournaments(context.Background())
	if err != nil {
		return nil, err
	}

	return scheduled, nil
}