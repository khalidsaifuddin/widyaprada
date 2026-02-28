package auth

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/pkg/auth"
)

// Logout menginvalidasi token dengan menambah ke blacklist
func (u *authUsecase) Logout(ctx context.Context, tokenString string) (*entity.LogoutResponse, error) {
	claims, err := auth.ParseToken(tokenString, u.cfg.JWTSecret)
	if err != nil {
		return nil, entity.ErrInvalidToken
	}

	var expiresAt time.Time
	if claims.ExpiresAt != nil {
		expiresAt = claims.ExpiresAt.Time
	}
	u.blacklist.Add(claims.ID, expiresAt)

	return &entity.LogoutResponse{
		Message: "Anda telah keluar.",
	}, nil
}
