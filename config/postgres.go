package configuration

import (
	"fmt"
	"goamartha/common"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"
)

var DB *gorm.DB

type Config struct {
	Host            string
	Port            int
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func ConfigDB(env Env) (*Config, error) {
	port, err := strconv.Atoi(env.Get("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	maxIdleConns, err := strconv.Atoi(env.Get("DB_MAX_IDLE_CONNS"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_MAX_IDLE_CONNS: %w", err)
	}

	maxOpenConns, err := strconv.Atoi(env.Get("DB_MAX_OPEN_CONNS"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_MAX_OPEN_CONNS: %w", err)
	}

	connMaxLifetime, err := time.ParseDuration(env.Get("DB_CONN_MAX_LIFETIME"))
	if err != nil {
		return nil, fmt.Errorf("invalid DB_CONN_MAX_LIFETIME: %w", err)
	}

	config := &Config{
		Host:            env.Get("DB_HOST"),
		Port:            port,
		User:            env.Get("DB_USER"),
		Password:        env.Get("DB_PASSWORD"),
		DBName:          env.Get("DB_NAME"),
		SSLMode:         env.Get("DB_SSLMODE"),
		MaxIdleConns:    maxIdleConns,
		MaxOpenConns:    maxOpenConns,
		ConnMaxLifetime: connMaxLifetime,
	}
	return config, nil
}
func ConnectDB(config *Config) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode)
	//dsn := os.Getenv("DATABASE_URL")

	var err error
	for retries := 0; retries < 10; retries++ { // create retry
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err == nil {
			sqlDB, err := db.DB()
			if err != nil {
				return fmt.Errorf("failed to get database instance: %w", err)
			}
			sqlDB.SetMaxIdleConns(config.MaxIdleConns)
			sqlDB.SetMaxOpenConns(config.MaxOpenConns)
			sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
			DB = db
			return nil
		}
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("failed to connect to database after multiple attempts: %w", err)

}

func NewDB(env Env) *gorm.DB {
	config, err := ConfigDB(env)
	if err != nil {
		common.Logger.Fatalf("failed to load config: %w", err)
		//return fmt.Errorf("failed to load config: %w", err)
	}
	if err := ConnectDB(config); err != nil {
		common.Logger.Fatalf("failed to connect to the database: %v", err)

	}
	return DB
}
