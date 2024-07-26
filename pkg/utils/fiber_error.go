package utils

import (
	"github.com/gofiber/fiber/v3"
)

func NewFiberError() fiber.Config {
	return fiber.Config{
		ErrorHandler: ErrorHandler,
	}
}
