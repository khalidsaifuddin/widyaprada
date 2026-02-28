package repository

import (
	"context"
	"time"
)

// PasswordResetToken domain (internal)
type PasswordResetToken struct {
	ID        string
	UserID    string
	TokenHash string
	ExpiresAt time.Time
}

// PasswordResetTokenRepo interface untuk operasi token reset password
type PasswordResetTokenRepo interface {
	Create(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error
	FindByTokenHash(ctx context.Context, tokenHash string) (*PasswordResetToken, error)
	Delete(ctx context.Context, id string) error
}
