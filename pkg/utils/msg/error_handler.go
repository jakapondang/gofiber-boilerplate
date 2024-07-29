package msg

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/sirupsen/logrus"
	"gofiber-boilerplatev3/internal/v1/interface/http/middlewares"
	"gofiber-boilerplatev3/pkg/utils/logruspack"
	"strconv"
)

type ErrorHandlerMSG struct {
	Code      int
	Message   string
	ErrorData any
}

func ErrorHandler(ctx fiber.Ctx, err error) error {
	var res ErrorHandlerMSG
	requestID := ctx.Locals(middlewares.RequestIDKey).(string)

	_, validationError := err.(ValidationError)
	_, badRequestError := err.(BadRequestError)
	_, notFoundError := err.(NotFoundError)
	_, unauthorizedError := err.(UnauthorizedError)

	res = ErrorHandlerMSG{
		Code:      fiber.StatusInternalServerError,
		Message:   "internal Error",
		ErrorData: err.Error(),
	}
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		if errJson != nil {
			logruspack.Logger.WithFields(logrus.Fields{
				"request_id": requestID,
				"code":       res.Code,
				"msg":        res.Message,
				"error":      errJson,
			}).Error("(" + strconv.Itoa(res.Code) + ") " + res.Message)
			return ctx.Status(res.Code).JSON(middlewares.ResponseError{
				Code:    res.Code,
				Message: res.Message,
				Error:   res.ErrorData,
			})
		}
		res = ErrorHandlerMSG{
			Code:      fiber.StatusBadRequest,
			Message:   "Bad Request",
			ErrorData: messages,
		}

	} else if badRequestError {
		res = ErrorHandlerMSG{
			Code:      fiber.StatusBadRequest,
			Message:   "Bad Request",
			ErrorData: err.Error(),
		}

	} else if notFoundError {
		res = ErrorHandlerMSG{
			Code:      fiber.StatusNotFound,
			Message:   "Not Found",
			ErrorData: err.Error(),
		}
	} else if unauthorizedError {
		res = ErrorHandlerMSG{
			Code:      fiber.StatusUnauthorized,
			Message:   "Unauthorized",
			ErrorData: err.Error(),
		}
	}
	logruspack.Logger.WithFields(logrus.Fields{
		"request_id": requestID,
		"code":       res.Code,
		"msg":        res.Message,
		"error":      res.ErrorData,
	}).Error("(" + strconv.Itoa(res.Code) + ") " + res.Message)
	return ctx.Status(res.Code).JSON(middlewares.ResponseError{
		Code:    res.Code,
		Message: res.Message,
		Error:   res.ErrorData,
	})
}

func ErrorHandler2(ctx fiber.Ctx, err error) error {
	var res middlewares.Response
	requestID := ctx.Locals(middlewares.RequestIDKey).(string)

	_, validationError := err.(ValidationError)
	if validationError {
		data := err.Error()
		var messages []map[string]interface{}

		errJson := json.Unmarshal([]byte(data), &messages)
		if errJson != nil {
			PanicLogging(errJson)
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(middlewares.Response{
			Code:    400,
			Message: "Bad Request",
			Data:    messages,
		})
	}

	_, badRequestError := err.(BadRequestError)
	if badRequestError {
		res = middlewares.Response{
			Code:    fiber.StatusBadRequest,
			Message: "Bad Request",
			Data:    err.Error(),
		}
		logruspack.Logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"msg":        res.Message,
			"error":      err.Error(),
		}).Error("(" + strconv.Itoa(fiber.StatusBadRequest) + ") " + res.Message)
		return ctx.Status(fiber.StatusBadRequest).JSON(res)
	}

	_, notFoundError := err.(NotFoundError)
	if notFoundError {
		res = middlewares.Response{
			Code:    fiber.StatusNotFound,
			Message: "Not Found",
			Data:    err.Error(),
		}

		return ctx.Status(fiber.StatusNotFound).JSON(res)
	}

	_, unauthorizedError := err.(UnauthorizedError)
	if unauthorizedError {
		res = middlewares.Response{
			Code:    401,
			Message: "Unauthorized",
			Data:    err.Error(),
		}

		return ctx.Status(fiber.StatusUnauthorized).JSON(res)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(middlewares.Response{
		Code:    500,
		Message: "internal Error",
		Data:    err.Error(),
	})

}
