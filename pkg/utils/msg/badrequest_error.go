package msg

import (
	"github.com/sirupsen/logrus"
	"gofiber-boilerplatev3/internal/v1/interface/http/middlewares"
	"gofiber-boilerplatev3/pkg/utils/logruspack"
)

type BadRequestError struct {
	Message   string
	RequestID string
}

func (badRequestError BadRequestError) Error() string {
	logruspack.Logger.WithFields(logrus.Fields{
		"request_id": middlewares.RequestID,
		"error":      badRequestError.Message,
	}).Error("Failed to unmarshal request body")
	return badRequestError.Message
}
