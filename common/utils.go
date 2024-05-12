package common

import (
	"championForge/common/types"

	"github.com/gofiber/fiber/v2"
)


func NewError(c *fiber.Ctx, code int, message string) error {
	response := types.Response{
		Message: message,
	}

	return c.Status(code).JSON(response)
}