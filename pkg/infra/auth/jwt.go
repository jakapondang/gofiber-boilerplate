package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/pkg/infra/config"
	"time"
)

type AccessTokenClaims struct {
	Data any `json:"data"`
	jwt.RegisteredClaims
}
type RefreshTokenClaims struct {
	ID  string `json:"id"`
	jti string `json:"jti"`
	jwt.RegisteredClaims
}

// Access Token
func GenerateAccessToken(config config.Config, data any) (string, error) {
	expirationTime := time.Now().Add(config.JWT.ExpAccessToken * time.Minute) // Access token expires in 15 minutes
	claims := &AccessTokenClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.JWT.AppName,
			Audience:  jwt.ClaimStrings{config.JWT.Audience},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Refresh Token
func GenerateRefreshToken(config config.Config, id string) (string, error) {

	expirationTime := time.Now().Add(config.JWT.ExpRefreshToken * time.Hour) // Refresh token expires in 7 days
	claims := &RefreshTokenClaims{
		ID:  id,
		jti: uuid.New().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.JWT.AppName,
			Audience:  jwt.ClaimStrings{config.JWT.Audience},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.JWT.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
