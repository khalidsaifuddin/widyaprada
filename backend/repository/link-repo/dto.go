package linkrepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// Link DTO
type Link struct {
	ID           string     `gorm:"column:id;primaryKey;type:uuid"`
	Title        string     `gorm:"column:title;size:255;not null"`
	URL          string     `gorm:"column:url;size:500;not null"`
	Description  string     `gorm:"column:description;type:text"`
	SortOrder    int        `gorm:"column:sort_order;default:0"`
	Status       string     `gorm:"column:status;size:20;default:Aktif"`
	OpenInNewTab bool       `gorm:"column:open_in_new_tab;default:false"`
	SatkerID     *string    `gorm:"column:satker_id;type:uuid"`
	CreatedAt    *time.Time `gorm:"column:created_at"`
	UpdatedAt    *time.Time `gorm:"column:updated_at"`
}

func (Link) TableName() string {
	return "links"
}

func (l *Link) ToListItem() entity.LinkListItem {
	createdAt := ""
	if l.CreatedAt != nil {
		createdAt = l.CreatedAt.UTC().Format(time.RFC3339)
	}
	return entity.LinkListItem{
		ID:           l.ID,
		Title:        l.Title,
		URL:          l.URL,
		Description:  l.Description,
		SortOrder:    l.SortOrder,
		Status:       l.Status,
		OpenInNewTab: l.OpenInNewTab,
		CreatedAt:    createdAt,
	}
}

func (l *Link) ToPublicItem() entity.LinkPublicItem {
	return entity.LinkPublicItem{
		ID:           l.ID,
		Title:        l.Title,
		URL:          l.URL,
		Description:  l.Description,
		SortOrder:    l.SortOrder,
		OpenInNewTab: l.OpenInNewTab,
	}
}
