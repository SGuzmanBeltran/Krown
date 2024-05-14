package utils

import (
	"championForge/common/types"

	"github.com/gofiber/fiber/v2"
)


func NewServiceResponse(code int, message string) *types.ServiceResponse {
	response := &types.ServiceResponse{
		Message: message,
		Status: code,
	}

	return response
}

func ServiceResponse(c *fiber.Ctx, r *types.ServiceResponse) error {
	return c.Status(r.Status).JSON(r)
}

func DirectResponse(c *fiber.Ctx, code int, message string) error {
	response := types.Response{
		Message: message,
	}
	return c.Status(code).JSON(response)
}

func CustomResponse(c *fiber.Ctx, code int, response any) error {
	return c.Status(code).JSON(response)
}