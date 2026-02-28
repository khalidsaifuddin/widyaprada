package userrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

// CreateUserWithRoles membuat user baru dan assign multiple roles
func (r *userRepo) CreateUserWithRoles(ctx context.Context, user *entity.User, roleIDs []string) error {
	if len(roleIDs) == 0 {
		return entity.ErrAtLeastOneRole
	}

	now := time.Now().UTC()
	userID := user.ID
	if userID == "" {
		userID = uuid.New().String()
	}

	u := User{}.FromEntity(*user)
	u.ID = userID
	u.CreatedAt = &now
	u.UpdatedAt = &now
	if err := r.db.WithContext(ctx).Create(&u).Error; err != nil {
		return err
	}

	for _, roleID := range roleIDs {
		ur := UserRole{
			UserID:    userID,
			RoleID:    roleID,
			CreatedAt: &now,
		}
		if err := r.db.WithContext(ctx).Create(&ur).Error; err != nil {
			return err
		}
	}
	return nil
}
