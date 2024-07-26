package users

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"goamartha/domain/model"
	"goamartha/exception"
)

type Controller struct {
	Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		Service: *service,
	}
}

// Find All Users
func (controller Controller) FindAll(c fiber.Ctx) error {

	result, err := controller.Service.FindAll(c.Context())
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	if err != nil {
		exception.PanicLogging(err)
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// Create User
func (controller Controller) Create(c fiber.Ctx) error {
	var request ModelRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		exception.PanicLogging(err)
	}

	result, err := controller.Service.Create(c.Context(), request)
	if err != nil {
		exception.PanicLogging(err)
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
