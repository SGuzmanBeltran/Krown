package user

import (
	"championForge/common/types"
	"championForge/config"
	"championForge/db"
	"championForge/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	userStore *Store
}

func NewUserService(userStore *Store) *UserService {
	return &UserService{userStore}
}

func (s *UserService) register(c *fiber.Ctx, payload db.CreateUserParams) *types.ServiceResponse {
	count, err := s.userStore.CheckUserByEmail(payload.Email)
	if err != nil {
		return utils.NewServiceResponse(fiber.StatusInternalServerError, fmt.Sprintf("Error getting the user by email, %s", err))
	}
	if count != 0 {
		return utils.NewServiceResponse(fiber.StatusConflict, "User with the same email already exists")
	}

	if err := s.userStore.CreateUser(c, payload); err != nil {
		return utils.NewServiceResponse(fiber.StatusConflict, "Error creating user")
	}
	return nil
}

func (s *UserService) login(c *fiber.Ctx, payload types.LoginUserPayload) (string, *types.ServiceResponse) {
	user, err := s.userStore.userQueries.GetUserByEmail(c.Context(), payload.Email)
	if err != nil {
		return "", utils.NewServiceResponse(fiber.StatusConflict, "Credentials not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return "", utils.NewServiceResponse(fiber.StatusConflict, "Credentials not found")
	}

	token := jwt.New(jwt.SigningMethodHS512)
	secretKey := []byte(config.Envs.SecretKey)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", utils.NewServiceResponse(fiber.StatusInternalServerError, "Error generating token")
	}
	return signedToken, nil
}
