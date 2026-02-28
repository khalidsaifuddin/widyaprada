package userrepo

import (
	"strings"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var defaultRoles = []struct {
	Code string
	Name string
}{
	{entity.DefaultRoleCodePeserta, "Peserta / Calon WP"},
	{"SUPER_ADMIN", "Super Admin"},
	{"ADMIN_UJIKOM", "Admin Uji Kompetensi"},
	{"ADMIN_SATKER", "Admin Satker"},
	{"VERIFIKATOR", "Verifikator"},
}

// SeedDefaultRoles ensures default roles exist (untuk registrasi & development)
func SeedDefaultRoles(db *gorm.DB) error {
	for _, r := range defaultRoles {
		var count int64
		if err := db.Model(&Role{}).Where("UPPER(code) = ?", strings.ToUpper(r.Code)).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		now := time.Now().UTC()
		role := Role{
			ID:        uuid.New().String(),
			Code:      r.Code,
			Name:      r.Name,
			CreatedAt: &now,
			UpdatedAt: &now,
		}
		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
