package journalrepo

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// mapJournalSort maps user-facing sort values (terbaru, terlama) to actual DB columns.
func mapJournalSort(sortBy, sortOrder string) (col, order string) {
	col = "published_at"
	order = "desc"
	sb := strings.ToLower(sortBy)
	so := strings.ToLower(sortOrder)
	switch sb {
	case "terbaru":
		col, order = "published_at", "desc"
	case "terlama":
		col, order = "published_at", "asc"
	case "published_at", "created_at", "updated_at", "title", "year":
		col = sb
		if so == "asc" || so == "desc" {
			order = so
		}
	}
	return col, order
}

func (r *journalRepo) ListPublished(ctx context.Context, req entity.GetJurnalListRequest) (*entity.GetJurnalListResponse, error) {
	var items []Journal
	db := r.db.WithContext(ctx).Model(&Journal{}).Where("deleted_at IS NULL AND status = ?", entity.JurnalStatusPublished)

	if req.Q != "" {
		q := "%" + req.Q + "%"
		db = db.Where("(title ILIKE ? OR abstract ILIKE ? OR content ILIKE ? OR author ILIKE ?)", q, q, q, q)
	}
	if req.Tahun != "" {
		if year, err := strconv.Atoi(req.Tahun); err == nil {
			db = db.Where("year = ? OR EXTRACT(YEAR FROM published_at) = ?", year, year)
		}
	}
	if req.Kategori != "" {
		db = db.Where("category = ?", req.Kategori)
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
	sortBy, sortOrder := mapJournalSort(req.Sort, req.SortOrder)
	order := sortBy + " " + sortOrder

	offset := (page - 1) * pageSize
	if err := db.Order(order + " NULLS LAST").Offset(int(offset)).Limit(int(pageSize)).Find(&items).Error; err != nil {
		return nil, err
	}

	list := make([]entity.JurnalListItem, len(items))
	for i := range items {
		list[i] = items[i].ToListItem()
	}

	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage++
	}

	return &entity.GetJurnalListResponse{
		Items:     list,
		TotalPage: totalPage,
		TotalData: total,
		Page:      page,
		PageSize:  pageSize,
	}, nil
}

func (r *journalRepo) GetByID(ctx context.Context, id string) (*entity.Jurnal, error) {
	var j Journal
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&j).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return toJournalEntity(&j), nil
}

func toJournalEntity(j *Journal) *entity.Jurnal {
	publishedAt, createdAt, updatedAt := "", "", ""
	if j.PublishedAt != nil {
		publishedAt = j.PublishedAt.UTC().Format(time.RFC3339)
	}
	if j.CreatedAt != nil {
		createdAt = j.CreatedAt.UTC().Format(time.RFC3339)
	}
	if j.UpdatedAt != nil {
		updatedAt = j.UpdatedAt.UTC().Format(time.RFC3339)
	}
	return &entity.Jurnal{
		ID:          j.ID,
		Title:       j.Title,
		Author:      j.Author,
		Abstract:    j.Abstract,
		Content:     j.Content,
		PdfURL:      j.PdfURL,
		PublishedAt: publishedAt,
		Status:      j.Status,
		Category:    j.Category,
		Year:        j.Year,
		UserID:      j.UserID,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func (r *journalRepo) ListPublishedForLanding(ctx context.Context, limit int) ([]entity.JurnalPublicItem, error) {
	var items []Journal
	err := r.db.WithContext(ctx).Model(&Journal{}).
		Where("deleted_at IS NULL AND status = ?", entity.JurnalStatusPublished).
		Order("published_at DESC NULLS LAST").
		Limit(limit).
		Find(&items).Error
	if err != nil {
		return nil, err
	}
	list := make([]entity.JurnalPublicItem, len(items))
	for i := range items {
		list[i] = items[i].ToPublicItem()
	}
	return list, nil
}

// Create for CMS (if needed later)
func (r *journalRepo) Create(ctx context.Context, j *entity.Jurnal) error {
	now := time.Now().UTC()
	var publishedAt *time.Time
	if j.PublishedAt != "" {
		if t, err := time.Parse(time.RFC3339, j.PublishedAt); err == nil {
			t := t.UTC()
			publishedAt = &t
		}
	}
	dto := Journal{
		ID:          uuid.New().String(),
		Title:       j.Title,
		Author:      j.Author,
		Abstract:    j.Abstract,
		Content:     j.Content,
		PdfURL:      j.PdfURL,
		PublishedAt: publishedAt,
		Status:      j.Status,
		Category:    j.Category,
		Year:        j.Year,
		UserID:      j.UserID,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	if err := r.db.WithContext(ctx).Create(&dto).Error; err != nil {
		return err
	}
	j.ID = dto.ID
	return nil
}

// Update journal by ID
func (r *journalRepo) Update(ctx context.Context, j *entity.Jurnal) error {
	var existing Journal
	if err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", j.ID).First(&existing).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	now := time.Now().UTC()
	var publishedAt *time.Time
	if j.PublishedAt != "" {
		if t, err := time.Parse(time.RFC3339, j.PublishedAt); err == nil {
			t := t.UTC()
			publishedAt = &t
		}
	}
	updates := map[string]interface{}{
		"title":        j.Title,
		"author":       j.Author,
		"abstract":     j.Abstract,
		"content":      j.Content,
		"pdf_url":      j.PdfURL,
		"published_at": publishedAt,
		"status":       j.Status,
		"category":     j.Category,
		"year":         j.Year,
		"updated_at":   &now,
	}
	return r.db.WithContext(ctx).Model(&Journal{}).Where("id = ?", j.ID).Updates(updates).Error
}
