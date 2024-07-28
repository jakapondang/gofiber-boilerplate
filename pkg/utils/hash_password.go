package utils

import (
	"gofiber-boilerplatev3/pkg/infra/config"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// HashPassword hashes a plain text password
func HashPassword(password string) (string, error) {
	costStr := config.AppConfig.Encryption.BcryptCost
	cost, err := strconv.Atoi(costStr)
	if err != nil {
		cost = bcrypt.DefaultCost
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with its plain text version
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
