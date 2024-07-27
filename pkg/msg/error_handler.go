package msg

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(ctx fiber.Ctx, err error) error {
	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		PanicLogging(errJson)

		return ctx.Status(fiber.StatusBadRequest).JSON(Response{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, badRequestError := err.(BadRequestError)
	if badRequestError {
		return ctx.Status(fiber.StatusBadRequest).JSON(Response{
			Code:    400,
			Message: "Bad Request",
			Data:    err.Error(),
		})
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		return ctx.Status(fiber.StatusNotFound).JSON(Response{
			Code:    404,
			Message: "Not Found",
			Data:    err.Error(),
		})
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		return ctx.Status(fiber.StatusUnauthorized).JSON(Response{
			Code:    401,
			Message: "Unauthorized",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(Response{
		Code:    500,
		Message: "internal Error",
		Data:    err.Error(),
	})

}