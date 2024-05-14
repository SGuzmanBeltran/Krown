package user

import (
	"context"
	"krown/common/types"
	"krown/db"
	"krown/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserStore interface {
	CheckUserByEmail(email string) (int64, error)
}

type Store struct {
	userQueries *db.Queries
}

func NewStore(userQueries *db.Queries) *Store {
	return &Store{userQueries: userQueries}
}

func (s *Store) CheckUserByEmail(email string) (int64, error) {
	count, err := s.userQueries.CheckUserByEmail(context.Background(), email)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (s *Store) CreateUser(c *fiber.Ctx, userParams db.CreateUserParams) *types.ServiceResponse {
	var hashed []byte
	hashed, err := bcrypt.GenerateFromPassword([]byte(userParams.Password), 8)
	if err != nil {
		return utils.NewServiceResponse(fiber.StatusInternalServerError, "Error hashing password")
	}
	userParams.Password = string(hashed)
	if err := s.userQueries.CreateUser(c.Context(), userParams); err != nil {
		return utils.NewServiceResponse(fiber.StatusInternalServerError, "Error creating user")
	}
	return nil
}
