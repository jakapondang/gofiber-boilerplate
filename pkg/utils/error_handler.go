package utils

import (
	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}
	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Code:    500,
		Message: "internal Error",
		Data:    err.Error(),
	})

}
