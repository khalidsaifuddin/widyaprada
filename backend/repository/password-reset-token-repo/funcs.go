package passwordresettokenrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *passwordResetTokenRepo) Create(ctx context.Context, userID, tokenHash string, expiresAt time.Time) error {
	now := time.Now().UTC()
	t := PasswordResetToken{
		ID:        uuid.New().String(),
		UserID:    userID,
		TokenHash: tokenHash,
		ExpiresAt: expiresAt,
		CreatedAt: &now,
	}
	return r.db.WithContext(ctx).Create(&t).Error
}

func (r *passwordResetTokenRepo) FindByTokenHash(ctx context.Context, tokenHash string) (*repository.PasswordResetToken, error) {
	var t PasswordResetToken
	err := r.db.WithContext(ctx).
		Where("token_hash = ? AND expires_at > ?", tokenHash, time.Now().UTC()).
		First(&t).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return t.ToEntity(), nil
}

func (r *passwordResetTokenRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&PasswordResetToken{}).Error
}
