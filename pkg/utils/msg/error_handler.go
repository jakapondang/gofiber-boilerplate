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
