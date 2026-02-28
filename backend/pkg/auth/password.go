package auth

import (
	"github.com/ProjectWidyaprada/backend/core/entity"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword menghasilkan bcrypt hash dari password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), entity.BcryptCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// VerifyPassword membandingkan password plain dengan hash
func VerifyPassword(plainPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
