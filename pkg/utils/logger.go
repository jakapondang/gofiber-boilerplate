package utils

import (
	"github.com/sirupsen/logrus"
	"gofiber-boilerplatev3/pkg/msg"
	"io"
	"os"
	"time"
)

var Logger *logrus.Logger

func NewLogger() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0770)
		msg.PanicLogging(err)
	}

	date := time.Now()
	logFile, err := os.OpenFile("logs/log_"+date.Format("01-02-2006_15")+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	msg.PanicLogging(err)
	if err == nil {
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		logger.SetOutput(multiWriter)
	}
	Logger = logger
}
