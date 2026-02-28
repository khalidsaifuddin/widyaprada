package linkrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *linkRepo) ListActiveForLanding(ctx context.Context, limit int) ([]entity.LinkPublicItem, error) {
	var items []Link
	err := r.db.WithContext(ctx).Model(&Link{}).
		Where("status = ?", entity.LinkStatusAktif).
		Order("sort_order ASC").
		Limit(limit).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	list := make([]entity.LinkPublicItem, len(items))
	for i := range items {
		list[i] = items[i].ToPublicItem()
	}
	return list, nil
}
