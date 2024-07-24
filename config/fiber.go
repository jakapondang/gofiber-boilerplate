package configuration

import (
	"github.com/gofiber/fiber/v3"
	"goamartha/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
