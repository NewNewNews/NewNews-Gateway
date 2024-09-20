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

// CustomClaims extends StandardClaims to include IsAdmin
type CustomClaims struct {
	jwt.StandardClaims
	IsAdmin bool `json:"is_admin"`
}

func (manager *JWTManager) Generate(userID string, isAdmin bool) (string, error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: time.Now().Add(manager.expiration).Unix(),
		},
		IsAdmin: isAdmin,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secretKey))
}

func (manager *JWTManager) Validate(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(manager.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

func (jm *JWTManager) GetExpiration() time.Duration {
	return jm.expiration
}
