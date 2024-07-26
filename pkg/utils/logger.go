package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func NewLogger() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.InfoLevel)
}
