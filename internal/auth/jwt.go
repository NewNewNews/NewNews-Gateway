package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	secretKey  string
	expiration time.Duration
}

func NewJWTManager(secretKey string, expiration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey, expiration: expiration}
}

func (manager *JWTManager) Generate(userID string, isAdmin bool) (string, error) {
	claims := jwt.StandardClaims{
		Subject:   userID,
		ExpiresAt: time.Now().Add(manager.expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}
