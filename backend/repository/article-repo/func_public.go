package articlerepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *articleRepo) ListPublished(ctx context.Context, req entity.GetArticleListRequest, limit int) ([]entity.ArticlePublicItem, int64, error) {
	var items []Article
	db := r.db.WithContext(ctx).Model(&Article{}).Where("deleted_at IS NULL AND status = ?", entity.ArticleStatusPublished)

	if req.Q != "" {
		q := "%" + req.Q + "%"
		db = db.Where("(title ILIKE ? OR excerpt ILIKE ? OR content ILIKE ?)", q, q, q)
	}
	if req.Category != "" {
		db = db.Where("category = ?", req.Category)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if limit < 1 {
		limit = 10
	}
	page, pageSize := req.Page, req.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	sortBy, sortOrder := mapArticleSort(req.SortBy, req.SortOrder)
	order := sortBy + " " + sortOrder

	offset := (page - 1) * pageSize
	if err := db.Order(order + " NULLS LAST").Offset(int(offset)).Limit(int(pageSize)).Find(&items).Error; err != nil {
		return nil, 0, err
	}

	list := make([]entity.ArticlePublicItem, len(items))
	for i := range items {
		list[i] = items[i].ToPublicItem()
	}
	return list, total, nil
}
