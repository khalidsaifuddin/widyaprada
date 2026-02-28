package userrepo

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// FindByEmail mencari user by email (case-insensitive)
func (r *userRepo) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var u User
	email = strings.TrimSpace(strings.ToLower(email))

	err := r.db.WithContext(ctx).
		Where("LOWER(email) = ?", email).
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
