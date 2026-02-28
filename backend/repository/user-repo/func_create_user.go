package userrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

// CreateUser membuat user baru dan assign role
func (r *userRepo) CreateUser(ctx context.Context, user *entity.User, roleID string) error {
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

	ur := UserRole{
		UserID:    userID,
		RoleID:    roleID,
		CreatedAt: &now,
	}
	return r.db.WithContext(ctx).Create(&ur).Error
}
