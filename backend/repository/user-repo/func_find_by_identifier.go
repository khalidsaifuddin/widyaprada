package userrepo

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// FindByEmailOrUsername mencari user by email atau username (identifier sudah dinormalisasi)
func (r *userRepo) FindByEmailOrUsername(ctx context.Context, identifier string) (*entity.User, error) {
	var u User
	identifier = strings.TrimSpace(strings.ToLower(identifier))

	err := r.db.WithContext(ctx).
		Where("LOWER(email) = ? OR LOWER(username) = ?", identifier, identifier).
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

// GetUserRoles mengambil roles user
func (r *userRepo) GetUserRoles(ctx context.Context, userID string) ([]entity.RoleInfo, error) {
	var roles []Role
	err := r.db.WithContext(ctx).
		Table("roles").
		Joins("INNER JOIN user_roles ON user_roles.role_id = roles.id").
		Where("user_roles.user_id = ?", userID).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}

	result := make([]entity.RoleInfo, 0, len(roles))
	for i := range roles {
		result = append(result, roles[i].ToEntity())
	}
	return result, nil
}
