package controller

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v3"
	"goamartha/domain/model"
	"goamartha/exception"
	"goamartha/service"
)

type BorrowerController struct {
	BorrowerService service.BorrowerService
	PaymentService  service.PaymentService
}

func NewBorrowerController(borrowerService *service.BorrowerService, paymentService *service.PaymentService) *BorrowerController {
	return &BorrowerController{
		BorrowerService: *borrowerService,
		PaymentService:  *paymentService,
	}
}

// Find All Borrowers
func (controller BorrowerController) FindAll(c fiber.Ctx) error {

	result, err := controller.BorrowerService.FindAll(c.Context())
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

// Get ID Borrower
func (controller BorrowerController) FindById(c fiber.Ctx) error {

	id := c.Params("id")

	result, err := controller.BorrowerService.FindById(c.Context(), id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	controller.PaymentService.PaymentSummaries(c.Context(), &result)
	err = controller.PaymentService.GetSchedule(c.Context(), &result)

	if err != nil {
		exception.PanicLogging(err)
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// Create Borrower
func (controller BorrowerController) Create(c fiber.Ctx) error {
	var request model.BorrowerRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		exception.PanicLogging(err)
	}

	user, _ := controller.BorrowerService.FindByUsername(c.Context(), request.Username)

	if user.Id > 0 && user.IsSettled == false {
		exception.PanicLogging(errors.New("This username account has an unpaid balance."))
	}

	result, err := controller.BorrowerService.Create(c.Context(), request)
	if err != nil {
		exception.PanicLogging(err)
	}

	controller.PaymentService.PaymentSummaries(c.Context(), &result)
	err = controller.PaymentService.GetSchedule(c.Context(), &result)

	if err != nil {
		exception.PanicLogging(err)
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

// Get ID Borrower
func (controller BorrowerController) UpdateTrxDate(c fiber.Ctx) error {
	id := c.Params("id")

	result, err := controller.BorrowerService.FindById(c.Context(), id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	controller.PaymentService.PaymentSummaries(c.Context(), &result)
	err = controller.PaymentService.GetSchedule(c.Context(), &result)
	var request model.BorrowerUpdateDateRequest

	if err := json.Unmarshal(c.Body(), &request); err != nil {
		exception.PanicLogging(err)
	}
	result.TransactionDate = request.TransactionDate

	err = controller.BorrowerService.Updates(c.Context(), &result)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
