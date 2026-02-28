package passwordresettokenrepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/repository"
)

// PasswordResetToken model untuk tabel password_reset_tokens
type PasswordResetToken struct {
	ID        string     `gorm:"column:id;primaryKey;type:uuid"`
	UserID    string     `gorm:"column:user_id;type:uuid;index;not null"`
	TokenHash string     `gorm:"column:token_hash;size:64;uniqueIndex;not null"`
	ExpiresAt time.Time  `gorm:"column:expires_at;not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
}

func (PasswordResetToken) TableName() string {
	return "password_reset_tokens"
}

func (t *PasswordResetToken) ToEntity() *repository.PasswordResetToken {
	if t == nil {
		return nil
	}
	return &repository.PasswordResetToken{
		ID:        t.ID,
		UserID:    t.UserID,
		TokenHash: t.TokenHash,
		ExpiresAt: t.ExpiresAt,
	}
}
