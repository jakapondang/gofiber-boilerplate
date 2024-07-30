package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

var (
	secret          = "your_secret"
	appName         = "your_app_name"
	audience        = "your_service_or_frontend_name"
	expAccessToken  = time.Minute * 15
	expRefreshToken = time.Hour * 24
)

func SetConfig(jwtSecret, jwtAppName, jwtAudience string, accessTokenExp, refreshTokenExp time.Duration) {
	secret = jwtSecret
	appName = jwtAppName
	audience = jwtAudience
	expAccessToken = accessTokenExp
	expRefreshToken = refreshTokenExp
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
	jti string `json:"jti"`
	jwt.RegisteredClaims
}

// GenerateAccessToken creates a new JWT access token
func GenerateAccessToken(data any) (string, error) {
	expirationTime := time.Now().Add(expAccessToken) // Access token expiration time
	claims := &AccessTokenClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    appName,
			Audience:  jwt.ClaimStrings{audience},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// GenerateRefreshToken creates a new JWT refresh token
func GenerateRefreshToken(id string) (string, error) {
	expirationTime := time.Now().Add(expRefreshToken) // Refresh token expiration time
	claims := &RefreshTokenClaims{
		ID:  id,
		jti: uuid.New().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    appName,
			Audience:  jwt.ClaimStrings{audience},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateAccessToken validates a given JWT access token
func ValidateAccessToken(tokenString string) (*AccessTokenClaims, error) {
	claims := &AccessTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	// Optionally check the audience and other claims here

	return claims, nil
}

// ValidateRefreshToken validates a given JWT refresh token
func ValidateRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	claims := &RefreshTokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	// Optionally check the audience and other claims here

	return claims, nil
}
