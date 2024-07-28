package msg

import (
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Send(c fiber.Ctx, code int, data interface{}) error {
	var res Response
	if code == 201 {
		res = Response{
			Code:    code,
			Message: "Successfully Created",
			Data:    data,
		}
		return c.Status(fiber.StatusCreated).JSON(res)
	} else {
		res = Response{
			Code:    code,
			Message: "Success",
			Data:    data,
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}

}
