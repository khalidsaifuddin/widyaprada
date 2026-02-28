package userrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// GetByID mengambil user by ID
func (r *userRepo) GetByID(ctx context.Context, userID string) (*entity.User, error) {
	var u User
	err := r.db.WithContext(ctx).
		Where("id = ?", userID).
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
