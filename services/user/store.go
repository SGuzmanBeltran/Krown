package user

import (
	"championForge/db"
	"context"
)

type Store struct {
	userQueries *db.Queries
}

func NewStore(userQueries *db.Queries) *Store {
	return &Store{userQueries: userQueries}
}

func (s *Store) GetUserByEmail(email string) (*db.User, error) {
	user, err := s.userQueries.GetUserByEmail(context.Background(), email)
	if err != nil {
        return nil, err
    }
    return &user, nil
}