package sliderepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

// ListPublishedForLanding returns slides with status Published, within date range
func (r *slideRepo) ListPublishedForLanding(ctx context.Context, limit int) ([]entity.SlidePublicItem, error) {
	now := time.Now().UTC()
	var items []Slide
	err := r.db.WithContext(ctx).Model(&Slide{}).
		Where("status = ?", entity.SlideStatusPublished).
		Where("(date_start IS NULL OR date_start <= ?)", now).
		Where("(date_end IS NULL OR date_end >= ?)", now).
		Order("sort_order ASC").
		Limit(limit).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	list := make([]entity.SlidePublicItem, len(items))
	for i := range items {
		list[i] = items[i].ToPublicItem()
	}
	return list, nil
}
