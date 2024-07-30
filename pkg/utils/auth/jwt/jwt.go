package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gofiber-boilerplatev3/pkg/infra/config"
	"time"
)

var (
	secret          string
	appName         string
	audience        string
	expAccessToken  time.Duration
	expRefreshToken time.Duration
)

func SetConfig(cfg config.JWTConfig) {
	secret = cfg.Secret
	appName = cfg.AppName
	audience = cfg.Audience
	expAccessToken = cfg.ExpAccessToken * time.Minute
	expRefreshToken = cfg.ExpRefreshToken * time.Hour
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenClaims struct {
	Data any `json:"data"`
	jwt.RegisteredClaims
}

type RefreshTokenClaims struct {
	ID  string `json:"id"`
	JTI string `json:"jti"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(data any) (string, error) {
	expirationTime := time.Now().Add(expAccessToken)
	claims := &AccessTokenClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    appName,
			Audience:  []string{audience},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateRefreshToken(id string) (string, error) {
	expirationTime := time.Now().Add(expRefreshToken)
	claims := &RefreshTokenClaims{
		ID:  id,
		JTI: uuid.New().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    appName,
			Audience:  []string{audience},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateAccessToken(tokenString string) (*AccessTokenClaims, error) {
	claims := &AccessTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				fmt.Println("Malformed token:", tokenString)
				return nil, errors.New("malformed token")
			case ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0:
				fmt.Println("Token is either expired or not active yet:", tokenString)
				return nil, errors.New("token is either expired or not active yet")
			default:
				fmt.Println("Token could not be handled:", tokenString)
				return nil, errors.New("could not handle this token")
			}
		} else {
			fmt.Println("Token parse error:", err)
			return nil, err
		}
	}

	if !token.Valid {
		fmt.Println("Token is invalid:", tokenString)
		return nil, errors.New("invalid token")
	}

	fmt.Println("Token is valid:", tokenString)
	return claims, nil
}

func ValidateRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	claims := &RefreshTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				fmt.Println("Malformed token:", tokenString)
				return nil, errors.New("malformed token")
			case ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0:
				fmt.Println("Token is either expired or not active yet:", tokenString)
				return nil, errors.New("token is either expired or not active yet")
			default:
				fmt.Println("Token could not be handled:", tokenString)
				return nil, errors.New("could not handle this token")
			}
		} else {
			fmt.Println("Token parse error:", err)
			return nil, err
		}
	}

	if !token.Valid {
		fmt.Println("Token is invalid:", tokenString)
		return nil, errors.New("invalid token")
	}

	fmt.Println("Token is valid:", tokenString)
	return claims, nil
}
