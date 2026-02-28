package userrepo

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// FindByUsername mencari user by username (case-insensitive)
func (r *userRepo) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	var u User
	username = strings.TrimSpace(strings.ToLower(username))

	err := r.db.WithContext(ctx).
		Where("LOWER(username) = ?", username).
		First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	user := u.ToEntity()
	return &user, nil
}
