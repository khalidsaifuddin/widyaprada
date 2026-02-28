package userrepo

import (
	"strings"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedDefaultRoles ensures PESERTA role exists (untuk registrasi)
func SeedDefaultRoles(db *gorm.DB) error {
	var count int64
	if err := db.Model(&Role{}).Where("UPPER(code) = ?", strings.ToUpper(entity.DefaultRoleCodePeserta)).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	now := time.Now().UTC()
	role := Role{
		ID:        uuid.New().String(),
		Code:      entity.DefaultRoleCodePeserta,
		Name:      "Peserta / Calon WP",
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	return db.Create(&role).Error
}
