package packagerepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// QuestionPackage DTO
type QuestionPackage struct {
	ID                 string         `gorm:"column:id;primaryKey;type:uuid"`
	Code               string         `gorm:"column:code;size:100;uniqueIndex;not null"`
	Name               string         `gorm:"column:name;size:255;not null"`
	Description        string         `gorm:"column:description;type:text"`
	Status             string         `gorm:"column:status;size:20;default:Draft"`
	VerificationStatus string         `gorm:"column:verification_status;size:20;default:Belum"`
	DeletedReason      string         `gorm:"column:deleted_reason"`
	CreatedAt          *time.Time     `gorm:"column:created_at"`
	UpdatedAt          *time.Time     `gorm:"column:updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (QuestionPackage) TableName() string {
	return "question_packages"
}

func (p *QuestionPackage) ToEntity() entity.PackageDetail {
	createdAt, updatedAt := "", ""
	if p.CreatedAt != nil {
		createdAt = p.CreatedAt.UTC().Format(time.RFC3339)
	}
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt.UTC().Format(time.RFC3339)
	}
	return entity.PackageDetail{
		ID:                 p.ID,
		Code:               p.Code,
		Name:               p.Name,
		Description:        p.Description,
		Status:             p.Status,
		VerificationStatus: p.VerificationStatus,
		DeletedReason:      p.DeletedReason,
		CreatedAt:          createdAt,
		UpdatedAt:          updatedAt,
	}
}

// PackageQuestionItem DTO
type PackageQuestionItem struct {
	PackageID  string     `gorm:"column:package_id;primaryKey;type:uuid"`
	QuestionID string     `gorm:"column:question_id;primaryKey;type:uuid"`
	SortOrder  int        `gorm:"column:sort_order;not null"`
	CreatedAt  *time.Time `gorm:"column:created_at"`
}

func (PackageQuestionItem) TableName() string {
	return "package_question_items"
}
