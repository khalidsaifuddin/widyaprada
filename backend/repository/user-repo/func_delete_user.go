package userrepo

import (
	"context"
	"time"
)

// DeleteUser soft delete user dengan alasan
func (r *userRepo) DeleteUser(ctx context.Context, userID, reason string) error {
	now := time.Now().UTC()
	return r.db.WithContext(ctx).Model(&User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"deleted_at":     now,
			"deleted_reason": reason,
		}).Error
}
