package controller

import (
	"github.com/gofiber/fiber/v3"
	"goamartha/domain/model"
	"goamartha/exception"
	"goamartha/service"
)

type PaymentController struct {
	BorrowerService service.BorrowerService
	PaymentService  service.PaymentService
}

func NewPaymentController(borrowerService *service.BorrowerService, paymentService *service.PaymentService) *PaymentController {
	return &PaymentController{
		BorrowerService: *borrowerService,
		PaymentService:  *paymentService,
	}
}

func (controller PaymentController) MakePayment(c fiber.Ctx) error {
	id := c.Params("id")

	borrow, err := controller.BorrowerService.FindById(c.Context(), id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}

	controller.PaymentService.PaymentSummaries(c.Context(), &borrow)
	err = controller.PaymentService.GetSchedule(c.Context(), &borrow)
	result, err := controller.PaymentService.MakePayment(c.Context(), &borrow)

	if err != nil {
		exception.PanicLogging(err)
	}
	if result.TotalPayment == borrow.PaymentSummary.OutstandingBallance { // Make Settled

		borrow.IsSettled = true

		controller.BorrowerService.Updates(c.Context(), &borrow)
	}
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}

func (controller PaymentController) MakeRePayment(c fiber.Ctx) error {
	id := c.Params("id")

	borrow, err := controller.BorrowerService.FindById(c.Context(), id)
	if err != nil {
		panic(exception.NotFoundError{
			Message: err.Error(),
		})
	}
	controller.PaymentService.PaymentSummaries(c.Context(), &borrow)
	err = controller.PaymentService.GetSchedule(c.Context(), &borrow)

	result, err := controller.PaymentService.MakeRePayment(c.Context(), &borrow)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    result,
	})
}
