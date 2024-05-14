package user

import (
	"championForge/common/types"
	"championForge/db"
	"championForge/utils"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userService *UserService
}

func NewHandler(userService *UserService) *Handler {
	return &Handler{userService}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Post("/register", h.handleRegister)
	router.Post("/login", h.handleLogin)
}

func (h *Handler) handleRegister(c *fiber.Ctx) error {
	var payload db.CreateUserParams
	if err := c.BodyParser(&payload); err != nil {
		return utils.DirectResponse(c, fiber.StatusBadRequest, "Error decoding payload")
	}

	err := h.userService.register(c, payload)
	if err != nil {
		return utils.ServiceResponse(c, err)
	}

	return utils.DirectResponse(c, fiber.StatusOK, "User created")
}

func (h *Handler) handleLogin(c *fiber.Ctx) error {
	var payload types.LoginUserPayload
	if err := c.BodyParser(&payload); err != nil {
		return utils.DirectResponse(c, fiber.StatusBadRequest, "Error decoding payload")
	}
	signedToken, err := h.userService.login(c, payload)
	if err != nil {
		return utils.ServiceResponse(c, err)
	}

	response := types.ResponseToken{
		Token: signedToken,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
