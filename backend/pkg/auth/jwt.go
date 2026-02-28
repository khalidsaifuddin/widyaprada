package auth

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Claims untuk JWT payload
type Claims struct {
	UserID   string   `json:"user_id"`
	Email    string   `json:"email"`
	Roles    []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

// GenerateToken membuat JWT access token
func GenerateToken(userID, email string, roles []string, secret string, expiryHours int) (string, int, error) {
	if secret == "" {
		secret = "default-secret"
	}
	if expiryHours <= 0 {
		expiryHours = 1
	}

	now := time.Now().UTC()
	exp := now.Add(time.Duration(expiryHours) * time.Hour)
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}

	expiresIn := expiryHours * 3600 // detik
	return tokenString, expiresIn, nil
}

// ParseToken mem-parse dan validasi JWT
func ParseToken(tokenString, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, entity.ErrInvalidToken
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, entity.ErrInvalidToken
	}
	return claims, nil
}
