package userrepo

import (
	"context"
)

func (r *userRepo) UpdatePassword(ctx context.Context, userID, passwordHash string) error {
	return r.db.WithContext(ctx).
		Model(&User{}).
		Where("id = ?", userID).
		Update("password_hash", passwordHash).Error
}
