package questionrepo

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// SeedDefaultCategories creates default question categories if none exist
func SeedDefaultCategories(db *gorm.DB) error {
	var count int64
	if err := db.Model(&QuestionCategory{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	now := time.Now().UTC()
	categories := []QuestionCategory{
		{ID: uuid.New().String(), Code: "UMUM", Name: "Umum", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Code: "TEORI", Name: "Teori", CreatedAt: &now, UpdatedAt: &now},
		{ID: uuid.New().String(), Code: "PRAKTIK", Name: "Praktik", CreatedAt: &now, UpdatedAt: &now},
	}
	return db.Create(&categories).Error
}
