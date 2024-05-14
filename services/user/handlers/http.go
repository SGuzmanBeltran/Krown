package handler

import (
	"krown/common/types"
	"krown/db"
	"krown/services/user"
	"krown/utils"

	"github.com/gofiber/fiber/v2"
)

type HttpHandler struct {
	userService *user.UserService
}

func NewHandler(userService *user.UserService) *HttpHandler {
	return &HttpHandler{userService}
}

func (h *HttpHandler) RegisterRoutes(router fiber.Router) {
	router.Post("/register", h.handleRegister)
	router.Post("/login", h.handleLogin)
}

func (h *HttpHandler) handleRegister(c *fiber.Ctx) error {
	var payload db.CreateUserParams
	if err := c.BodyParser(&payload); err != nil {
		return utils.DirectResponse(c, fiber.StatusBadRequest, "Error decoding payload")
	}

	err := h.userService.Register(c, payload)
	if err != nil {
		return utils.ServiceResponse(c, err)
	}

	return utils.DirectResponse(c, fiber.StatusOK, "User created")
}

func (h *HttpHandler) handleLogin(c *fiber.Ctx) error {
	var payload types.LoginUserPayload
	if err := c.BodyParser(&payload); err != nil {
		return utils.DirectResponse(c, fiber.StatusBadRequest, "Error decoding payload")
	}
	signedToken, err := h.userService.Login(c, payload)
	if err != nil {
		return utils.ServiceResponse(c, err)
	}

	response := types.ResponseToken{
		Token: signedToken,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
