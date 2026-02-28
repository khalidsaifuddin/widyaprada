package userrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// UpdateUser mengupdate user dan sync roles
func (r *userRepo) UpdateUser(ctx context.Context, user *entity.User, roleIDs []string) error {
	updates := map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"username":   user.Username,
		"nip":        user.NIP,
		"satker_id":  user.SatkerID,
		"is_active":  user.IsActive,
		"updated_at": time.Now().UTC(),
	}

	if err := r.db.WithContext(ctx).Model(&User{}).Where("id = ?", user.ID).Updates(updates).Error; err != nil {
		return err
	}

	// Sync password if provided (caller hashes it)
	if user.PasswordHash != "" {
		if err := r.UpdatePassword(ctx, user.ID, user.PasswordHash); err != nil {
			return err
		}
	}

	// Sync roles
	if roleIDs != nil {
		// Delete existing user_roles
		if err := r.db.WithContext(ctx).Where("user_id = ?", user.ID).Delete(&UserRole{}).Error; err != nil {
			return err
		}
		// Insert new roles
		now := time.Now().UTC()
		for _, roleID := range roleIDs {
			if roleID == "" {
				continue
			}
			ur := UserRole{UserID: user.ID, RoleID: roleID, CreatedAt: &now}
			if err := r.db.WithContext(ctx).Create(&ur).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
