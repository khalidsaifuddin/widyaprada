package linkrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *linkRepo) List(ctx context.Context, req entity.GetLinkListRequest, satkerFilter *string) (*entity.GetLinkListResponse, error) {
	var items []Link
	db := r.db.WithContext(ctx).Model(&Link{})

	if req.Status != "" {
		db = db.Where("status = ?", req.Status)
	}
	if req.SatkerID != "" {
		db = db.Where("satker_id = ?", req.SatkerID)
	}
	if satkerFilter != nil && *satkerFilter != "" {
		db = db.Where("satker_id = ?", *satkerFilter)
	}

	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, err
	}

	page, pageSize := req.Page, req.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}
	sortBy := req.SortBy
	if sortBy == "" {
		sortBy = "sort_order"
	}
	sortOrder := req.SortOrder
	if sortOrder == "" {
		sortOrder = "asc"
	}
	order := sortBy + " " + sortOrder

	offset := (page - 1) * pageSize
	if err := db.Order(order).Offset(int(offset)).Limit(int(pageSize)).Find(&items).Error; err != nil {
		return nil, err
	}

	list := make([]entity.LinkListItem, len(items))
	for i := range items {
		list[i] = items[i].ToListItem()
	}

	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage++
	}

	return &entity.GetLinkListResponse{
		Items:     list,
		TotalPage: totalPage,
		TotalData: total,
		Page:      page,
		PageSize:  pageSize,
	}, nil
}

func (r *linkRepo) GetByID(ctx context.Context, id string) (*entity.Link, error) {
	var l Link
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&l).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	createdAt, updatedAt := "", ""
	if l.CreatedAt != nil {
		createdAt = l.CreatedAt.UTC().Format(time.RFC3339)
	}
	if l.UpdatedAt != nil {
		updatedAt = l.UpdatedAt.UTC().Format(time.RFC3339)
	}
	return &entity.Link{
		ID:           l.ID,
		Title:        l.Title,
		URL:          l.URL,
		Description:  l.Description,
		SortOrder:    l.SortOrder,
		Status:       l.Status,
		OpenInNewTab: l.OpenInNewTab,
		SatkerID:     l.SatkerID,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}, nil
}

func (r *linkRepo) Create(ctx context.Context, l *entity.Link) (string, error) {
	now := time.Now().UTC()
	id := uuid.New().String()
	dto := Link{
		ID:           id,
		Title:        l.Title,
		URL:          l.URL,
		Description:  l.Description,
		SortOrder:    l.SortOrder,
		Status:       l.Status,
		OpenInNewTab: l.OpenInNewTab,
		SatkerID:     l.SatkerID,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	if err := r.db.WithContext(ctx).Create(&dto).Error; err != nil {
		return "", err
	}
	return id, nil
}

func (r *linkRepo) Update(ctx context.Context, l *entity.Link) error {
	upd := map[string]interface{}{
		"title":          l.Title,
		"url":            l.URL,
		"description":    l.Description,
		"sort_order":     l.SortOrder,
		"status":         l.Status,
		"open_in_new_tab": l.OpenInNewTab,
		"updated_at":     time.Now().UTC(),
	}
	return r.db.WithContext(ctx).Model(&Link{}).Where("id = ?", l.ID).Updates(upd).Error
}

func (r *linkRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Link{}).Error
}
