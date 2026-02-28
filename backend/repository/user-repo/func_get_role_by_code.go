package userrepo

import (
	"context"
	"strings"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// GetRoleByCode mencari role by code (case-insensitive)
func (r *userRepo) GetRoleByCode(ctx context.Context, code string) (*entity.RoleInfo, error) {
	var role Role
	err := r.db.WithContext(ctx).
		Where("UPPER(code) = ?", strings.ToUpper(code)).
		First(&role).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	ri := role.ToEntity()
	return &ri, nil
}
