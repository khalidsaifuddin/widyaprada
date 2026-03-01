package journalrepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// Journal DTO
type Journal struct {
	ID          string         `gorm:"column:id;primaryKey;type:uuid"`
	Title       string         `gorm:"column:title;size:500;not null"`
	Author      string         `gorm:"column:author;size:255"`
	Abstract    string         `gorm:"column:abstract;type:text"`
	Content     string         `gorm:"column:content;type:text"`
	PdfURL      string         `gorm:"column:pdf_url;size:1000"`
	PublishedAt *time.Time     `gorm:"column:published_at"`
	Status      string         `gorm:"column:status;size:50;default:Draft"`
	Category    string         `gorm:"column:category;size:100"`
	Year        int            `gorm:"column:year"`
	UserID      *string        `gorm:"column:user_id;type:uuid;index"`
	SubmittedAt *time.Time     `gorm:"column:submitted_at"`
	CreatedAt   *time.Time     `gorm:"column:created_at"`
	UpdatedAt   *time.Time     `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Journal) TableName() string {
	return "journals"
}

func (j *Journal) ToListItem() entity.JurnalListItem {
	publishedAt := ""
	if j.PublishedAt != nil {
		publishedAt = j.PublishedAt.UTC().Format(time.RFC3339)
	}
	return entity.JurnalListItem{
		ID:          j.ID,
		Title:       j.Title,
		Author:      j.Author,
		Abstract:    j.Abstract,
		PublishedAt: publishedAt,
		Year:        j.Year,
		Category:    j.Category,
	}
}

func (j *Journal) ToPublicItem() entity.JurnalPublicItem {
	publishedAt := ""
	if j.PublishedAt != nil {
		publishedAt = j.PublishedAt.UTC().Format(time.RFC3339)
	}
	return entity.JurnalPublicItem{
		ID:          j.ID,
		Title:       j.Title,
		Author:      j.Author,
		Abstract:    j.Abstract,
		PublishedAt: publishedAt,
	}
}
