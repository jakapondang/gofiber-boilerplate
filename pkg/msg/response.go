package msg

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Send(c fiber.Ctx, data interface{}) error {
	fmt.Println(c.Status)
	var response Response
	//if code == 201 {
	//	response = Response{
	//		Code:    code,
	//		Message: "Success Created",
	//		Data:    data,
	//	}
	//} else {
	//	response = Response{
	//		Code:    code,
	//		Message: "Success",
	//		Data:    data,
	//	}
	//}
	response = Response{
		Code:    200,
		Message: "Success",
		Data:    data,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
