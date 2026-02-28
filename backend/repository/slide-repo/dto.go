package sliderepo

import (
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// Slide DTO
type Slide struct {
	ID        string     `gorm:"column:id;primaryKey;type:uuid"`
	ImageURL  string     `gorm:"column:image_url;size:500;not null"`
	Title     string     `gorm:"column:title;size:255"`
	Subtitle  string     `gorm:"column:subtitle;size:500"`
	LinkURL   string     `gorm:"column:link_url;size:500"`
	CTALabel  string     `gorm:"column:cta_label;size:100"`
	SortOrder int        `gorm:"column:sort_order;default:0"`
	Status    string     `gorm:"column:status;size:20;default:Draft"`
	DateStart *time.Time `gorm:"column:date_start"`
	DateEnd   *time.Time `gorm:"column:date_end"`
	SatkerID  *string    `gorm:"column:satker_id;type:uuid"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (Slide) TableName() string {
	return "slides"
}

func (s *Slide) ToListItem() entity.SlideListItem {
	dateStart, dateEnd, createdAt := "", "", ""
	if s.DateStart != nil {
		dateStart = s.DateStart.UTC().Format(time.RFC3339)
	}
	if s.DateEnd != nil {
		dateEnd = s.DateEnd.UTC().Format(time.RFC3339)
	}
	if s.CreatedAt != nil {
		createdAt = s.CreatedAt.UTC().Format(time.RFC3339)
	}
	return entity.SlideListItem{
		ID:        s.ID,
		ImageURL:  s.ImageURL,
		Title:     s.Title,
		Subtitle:  s.Subtitle,
		LinkURL:   s.LinkURL,
		CTALabel:  s.CTALabel,
		SortOrder: s.SortOrder,
		Status:    s.Status,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		CreatedAt: createdAt,
	}
}

func (s *Slide) ToPublicItem() entity.SlidePublicItem {
	return entity.SlidePublicItem{
		ID:        s.ID,
		ImageURL:  s.ImageURL,
		Title:     s.Title,
		Subtitle:  s.Subtitle,
		LinkURL:   s.LinkURL,
		CTALabel:  s.CTALabel,
		SortOrder: s.SortOrder,
	}
}
