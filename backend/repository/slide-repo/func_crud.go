package sliderepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (r *slideRepo) List(ctx context.Context, req entity.GetSlideListRequest, satkerFilter *string) (*entity.GetSlideListResponse, error) {
	var items []Slide
	db := r.db.WithContext(ctx).Model(&Slide{})

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

	list := make([]entity.SlideListItem, len(items))
	for i := range items {
		list[i] = items[i].ToListItem()
	}

	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage++
	}

	return &entity.GetSlideListResponse{
		Items:     list,
		TotalPage: totalPage,
		TotalData: total,
		Page:      page,
		PageSize:  pageSize,
	}, nil
}

func (r *slideRepo) GetByID(ctx context.Context, id string) (*entity.Slide, error) {
	var s Slide
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&s).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	dateStart, dateEnd, createdAt, updatedAt := "", "", "", ""
	if s.DateStart != nil {
		dateStart = s.DateStart.UTC().Format(time.RFC3339)
	}
	if s.DateEnd != nil {
		dateEnd = s.DateEnd.UTC().Format(time.RFC3339)
	}
	if s.CreatedAt != nil {
		createdAt = s.CreatedAt.UTC().Format(time.RFC3339)
	}
	if s.UpdatedAt != nil {
		updatedAt = s.UpdatedAt.UTC().Format(time.RFC3339)
	}
	return &entity.Slide{
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
		SatkerID:  s.SatkerID,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (r *slideRepo) Create(ctx context.Context, s *entity.Slide) error {
	now := time.Now().UTC()
	id := s.ID
	if id == "" {
		id = uuid.New().String()
	}
	var dateStart, dateEnd *time.Time
	if s.DateStart != "" {
		if t, err := time.Parse(time.RFC3339, s.DateStart); err == nil {
			t := t.UTC()
			dateStart = &t
		}
	}
	if s.DateEnd != "" {
		if t, err := time.Parse(time.RFC3339, s.DateEnd); err == nil {
			t := t.UTC()
			dateEnd = &t
		}
	}
	dto := Slide{
		ID:        id,
		ImageURL:  s.ImageURL,
		Title:     s.Title,
		Subtitle:  s.Subtitle,
		LinkURL:   s.LinkURL,
		CTALabel:  s.CTALabel,
		SortOrder: s.SortOrder,
		Status:    s.Status,
		DateStart: dateStart,
		DateEnd:   dateEnd,
		SatkerID:  s.SatkerID,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
	return r.db.WithContext(ctx).Create(&dto).Error
}

func (r *slideRepo) Update(ctx context.Context, s *entity.Slide) error {
	var dateStart, dateEnd *time.Time
	if s.DateStart != "" {
		if t, err := time.Parse(time.RFC3339, s.DateStart); err == nil {
			t := t.UTC()
			dateStart = &t
		}
	}
	if s.DateEnd != "" {
		if t, err := time.Parse(time.RFC3339, s.DateEnd); err == nil {
			t := t.UTC()
			dateEnd = &t
		}
	}
	upd := map[string]interface{}{
		"image_url":  s.ImageURL,
		"title":      s.Title,
		"subtitle":   s.Subtitle,
		"link_url":   s.LinkURL,
		"cta_label":  s.CTALabel,
		"sort_order": s.SortOrder,
		"status":     s.Status,
		"date_start": dateStart,
		"date_end":   dateEnd,
		"updated_at": time.Now().UTC(),
	}
	return r.db.WithContext(ctx).Model(&Slide{}).Where("id = ?", s.ID).Updates(upd).Error
}

func (r *slideRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Slide{}).Error
}
