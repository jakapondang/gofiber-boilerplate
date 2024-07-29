package logruspack

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"time"
)

type CustomLog struct {
	Logger *logrus.Logger
}

func (l *CustomLog) LogMode(level logger.LogLevel) logger.Interface {
	return &CustomLog{
		Logger: l.Logger,
	}
}

func (l *CustomLog) Info(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Infof(msg, args...)
}

func (l *CustomLog) Warn(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Warnf(msg, args...)
}

func (l *CustomLog) Error(ctx context.Context, msg string, args ...interface{}) {
	l.Logger.Errorf(msg, args...)
}

func (l *CustomLog) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	duration := time.Since(begin)
	sql, rows := fc()
	if err != nil {
		l.Logger.Errorf("SQL Error: %v | Duration: %v | SQL: %s | Rows: %d", err, duration, sql, rows)
	} else {
		l.Logger.Debugf("SQL Duration: %v | SQL: %s | Rows: %d", duration, sql, rows)
	}
}
