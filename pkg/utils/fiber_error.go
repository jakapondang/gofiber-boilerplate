package utils

import (
	"github.com/gofiber/fiber/v3"
	"gofiber-boilerplatev3/pkg/utils/msg"
)

func NewFiberError() fiber.Config {
	return fiber.Config{
		ErrorHandler: msg.ErrorHandler,
		//Prefork:      false,
		ServerHeader: "Fiber",
		AppName:      "FiberApp",
	}
}
