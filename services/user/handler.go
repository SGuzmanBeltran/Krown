package user

import (
	"championForge/common"
	"championForge/common/types"
	"championForge/config"
	"championForge/db"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	userStore Store
}

func NewHandler(userStore Store) *Handler {
	return &Handler{userStore}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/register", h.handleRegister)
	router.Post("/login", h.handleLogin)
	router.Get("/", h.getUsers)
}

func (h *Handler) handleRegister(c *fiber.Ctx) error {
	var payload db.CreateUserParams
	if err := c.BodyParser(&payload); err != nil {
		return common.NewError(c, fiber.StatusBadRequest, "Error decoding payload")
	}
	count, err := h.userStore.CheckUserByEmail(payload.Email)
	if err != nil {
		return common.NewError(c, fiber.StatusInternalServerError, fmt.Sprintf("Error getting the user by email, %s", err))
	}
	if count != 0 {
		return common.NewError(c, fiber.StatusConflict, "User with the same email already exists")
	}

	if err = h.userStore.CreateUser(c, payload); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) handleLogin(c *fiber.Ctx) error {
	var payload types.LoginUserPayload
	if err := c.BodyParser(&payload); err != nil {
		return common.NewError(c, fiber.StatusBadRequest, "Error decoding payload")
	}

	user, err := h.userStore.userQueries.GetUserByEmail(c.Context(), payload.Email)
	if err != nil {
		return common.NewError(c, fiber.StatusConflict, "Credentials not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return common.NewError(c, fiber.StatusConflict, "Credentials not found")
	}
	token := jwt.New(jwt.SigningMethodHS512)
	secretKey := []byte(config.Envs.SecretKey)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return common.NewError(c, fiber.StatusInternalServerError, "Error generating token")
	}
	//todo: send token
	response := types.ResponseToken{
		Token: signedToken,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h *Handler) getUsers(c *fiber.Ctx) error {
	fmt.Println("users")
	return c.SendStatus(fiber.StatusOK)
}
