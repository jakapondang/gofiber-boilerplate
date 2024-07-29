package database

import (
	"fmt"
	"gofiber-boilerplatev3/pkg/infra/config"
	"gofiber-boilerplatev3/pkg/infra/middleware/logruspack"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect(cfg config.Config) error {
	var err error
	dsn := cfg.Database

	logger := logruspack.Logger // Access Logrus logger instance

	retryCount := 10
	retryDelay := 2 * time.Second

	for retries := 0; retries < retryCount; retries++ {
		db, err := gorm.Open(postgres.Open(dsn.Url), &gorm.Config{
			Logger: &logruspack.CustomLog{Logger: logger}, // Use your custom logger
		})
		if err == nil {
			sqlDB, err := db.DB()
			if err != nil {
				return fmt.Errorf("failed to get database instance: %w", err)
			}
			sqlDB.SetMaxIdleConns(dsn.MaxIdleConns)
			sqlDB.SetMaxOpenConns(dsn.MaxOpenConns)
			sqlDB.SetConnMaxLifetime(dsn.ConnMaxLifetime)
			DB = db
			logger.Info("Database connection established successfully.")
			return nil
		}

		logger.Warnf("Database connection failed (attempt %d/%d): %v", retries+1, retryCount, err)
		time.Sleep(retryDelay)
	}

	return fmt.Errorf("failed to connect to database after %d attempts: %w", retryCount, err)
}
