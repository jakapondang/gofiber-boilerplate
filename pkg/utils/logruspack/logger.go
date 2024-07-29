package logruspack

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"time"
)

const LogsRootDir = "logs"
const LogsFileName = "app"

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

func Init() {
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
	if _, err := os.Stat(LogsRootDir); os.IsNotExist(err) {
		err := os.Mkdir(LogsRootDir, 0755) // Adjust permissions as needed
		if err != nil {
			logger.Fatalf("Failed to create logs directory: %v", err)
		}
	}

	// Set up log file rotation
	date := time.Now()
	logFile := &lumberjack.Logger{
		Filename:   LogsRootDir + "/" + LogsFileName + "_" + date.Format("01-02-2006") + ".log",
		MaxSize:    10, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
		Compress:   true,
	}

	// Set up JSON formatter for file output
	jsonFormatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@time",
			logrus.FieldKeyMsg:  "msg",
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
