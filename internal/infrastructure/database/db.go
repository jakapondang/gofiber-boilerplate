package database

import (
	"fmt"
	"gofiber-boilerplatev3/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var DB *gorm.DB

// Connect initializes the database connection
func Connect(config config.Config) error {
	var err error
	dsn := config.Database

	for retries := 0; retries < 10; retries++ { // create retry
		db, err := gorm.Open(postgres.Open(dsn.Url), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
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
			return nil
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("failed to connect to database after multiple attempts: %w", err)
}
