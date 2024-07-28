package config

import (
	"github.com/gofiber/fiber/v3/log"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Database   DatabaseConfig   `mapstructure:"database"`
	Log        LogConfig        `mapstructure:"log"`
	Encryption EncryptionConfig `mapstructure:"encryption"`
}

var AppConfig Config

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type DatabaseConfig struct {
	//Host            string        `mapstructure:"host"`
	//Port            int           `mapstructure:"port"`
	//User            string        `mapstructure:"user"`
	//Password        string        `mapstructure:"password"`
	//Name            string        `mapstructure:"name"`
	Url             string        `mapstructure:"url"`
	SSLMode         string        `mapstructure:"sslmode"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

type EncryptionConfig struct {
	BcryptCost string `mapstructure:"bcrypt_cost"`
	JwtSecret  string `mapstructure:"jwt_secret"`
}

func LoadConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}

func init() {
	if err := LoadConfig(); err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
}
