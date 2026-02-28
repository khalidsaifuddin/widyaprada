package examplerepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// Example is the DB model for examples table
type Example struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string         `gorm:"column:name;size:255"`
	CreatedAt *time.Time     `gorm:"column:created_at"`
	UpdatedAt *time.Time     `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Example) TableName() string {
	return "examples"
}

func (e *Example) ToEntity() entity.Example {
	return entity.Example{
		ID:        e.ID,
		Name:      e.Name,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
