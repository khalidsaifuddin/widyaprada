package articlerepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"gorm.io/gorm"
)

// Article DTO
type Article struct {
	ID           string     `gorm:"column:id;primaryKey;type:uuid"`
	Title        string     `gorm:"column:title;size:500;not null"`
	Slug         string     `gorm:"column:slug;size:300;uniqueIndex;not null"`
	Content      string     `gorm:"column:content;type:text"`
	Excerpt      string     `gorm:"column:excerpt;type:text"`
	ThumbnailURL string     `gorm:"column:thumbnail_url;size:500"`
	PublishedAt  *time.Time `gorm:"column:published_at"`
	Status       string     `gorm:"column:status;size:20;default:Draft"`
	AuthorName   string     `gorm:"column:author_name;size:255"`
	Category     string     `gorm:"column:category;size:100"`
	SatkerID     *string    `gorm:"column:satker_id;type:uuid"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
}

func (Article) TableName() string {
	return "articles"
}

func (a *Article) ToListItem() entity.ArticleListItem {
	publishedAt, createdAt := "", ""
	if a.PublishedAt != nil {
		publishedAt = a.PublishedAt.UTC().Format(time.RFC3339)
	}
	if a.CreatedAt != nil {
		createdAt = a.CreatedAt.UTC().Format(time.RFC3339)
	}
	return entity.ArticleListItem{
		ID:           a.ID,
		Title:        a.Title,
		Slug:         a.Slug,
		Excerpt:      a.Excerpt,
		ThumbnailURL: a.ThumbnailURL,
		PublishedAt:  publishedAt,
		Status:       a.Status,
		AuthorName:   a.AuthorName,
		Category:     a.Category,
		CreatedAt:    createdAt,
	}
}

func (a *Article) ToPublicItem() entity.ArticlePublicItem {
	publishedAt := ""
	if a.PublishedAt != nil {
		publishedAt = a.PublishedAt.UTC().Format(time.RFC3339)
	}
	return entity.ArticlePublicItem{
		ID:           a.ID,
		Title:        a.Title,
		Slug:         a.Slug,
		Excerpt:      a.Excerpt,
		ThumbnailURL: a.ThumbnailURL,
		PublishedAt:  publishedAt,
	}
}
