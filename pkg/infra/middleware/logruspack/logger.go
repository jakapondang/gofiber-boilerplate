package logruspack

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
)

var Logger *logrus.Logger

type LogFileHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
}

func (hook *LogFileHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

func (hook *LogFileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func New() {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Set up Text formatter for stdout output
	textFormatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
		PadLevelText:    true,
		DisableQuote:    true,
	}

	// Set the formatter for stdout
	logger.SetFormatter(textFormatter)

	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755) // Adjust permissions as needed
		if err != nil {
			logger.Fatalf("Failed to create logs directory: %v", err)
		}
	}

	// Set up log file rotation
	logFile := &lumberjack.Logger{
		Filename:   "logs/logfile.log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}

	// Set up JSON formatter for file output
	jsonFormatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	}

	// Add hook for logging to file with JSON format
	logger.AddHook(&LogFileHook{
		Writer:    logFile,
		Formatter: jsonFormatter,
	})

	// Also log to stdout
	logger.SetOutput(os.Stdout)

	Logger = logger
}
