package middlewares

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gofiber-boilerplatev3/pkg/utils/logruspack"

	"time"
)

const RequestIDKey = "X-Request-ID"

func RequestID() fiber.Handler {
	return func(c fiber.Ctx) error {
		requestID := c.Get(RequestIDKey)
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Locals(RequestIDKey, requestID)
		c.Set(RequestIDKey, requestID)
		return c.Next()
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func Send(c fiber.Ctx, code int, data interface{}) error {
	var res Response
	var message string
	if code == fiber.StatusCreated {
		message = "Successfully Created"
	} else {
		message = "Success"
	}
	res = Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	if code == fiber.StatusCreated {
		return c.Status(fiber.StatusCreated).JSON(res)
	} else {
		return c.Status(fiber.StatusOK).JSON(res)
	}
}

// Log Request Response
func LogRequestResponse() fiber.Handler {
	return func(c fiber.Ctx) error {
		// Extract request ID
		requestID := c.Locals(RequestIDKey).(string)

		// Log the request
		reqBody := c.Body()
		reqHeaders := c.GetReqHeaders()
		logruspack.Logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     c.Method(),
			"url":        c.OriginalURL(),
			"headers":    reqHeaders,
			"body":       string(reqBody),
		}).Info("Request")

		// Process the request
		start := time.Now()
		err := c.Next()
		duration := time.Since(start)

		// Log the response
		resBody := c.Response().Body()
		resHeaders := c.Response().Header.String()

		logruspack.Logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"status":     c.Response().StatusCode(),
			"duration":   duration.String(),
			"headers":    resHeaders,
			"body":       string(resBody),
		}).Info("Response")

		return err
	}
}
