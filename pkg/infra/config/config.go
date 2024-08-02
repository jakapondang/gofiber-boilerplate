package config

import (
	"github.com/spf13/viper"
	"gofiber-boilerplatev3/pkg/utils/logruspack"
	"time"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Database   DatabaseConfig   `mapstructure:"database"`
	Encryption EncryptionConfig `mapstructure:"encryption"`
	JWT        JWTConfig        `mapstructure:"jwt"`
	Mail       MailConfig       `mapstructure:"mail"`
}

var AppConfig Config

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Url             string        `mapstructure:"url"`
	SSLMode         string        `mapstructure:"sslmode"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

type EncryptionConfig struct {
	BcryptCost int `mapstructure:"bcrypt_cost"`
}

type JWTConfig struct {
	Secret          string        `mapstructure:"secret"`
	AppName         string        `mapstructure:"app_name"`
	Audience        string        `mapstructure:"audience"`
	ExpAccessToken  time.Duration `mapstructure:"exp_access_token"`
	ExpRefreshToken time.Duration `mapstructure:"exp_refresh_token"`
}
type MailConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// other configurations
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

func Init() {
	if err := LoadConfig(); err != nil {
		logruspack.Logger.Fatalf("Error loading config file: %v", err)
	}
}
