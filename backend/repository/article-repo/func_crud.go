package articlerepo

import (
	"context"
	"strings"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// mapArticleSort maps user-facing sort values (terbaru, terlama) to actual DB columns.
func mapArticleSort(sortBy, sortOrder string) (col, order string) {
	col = "published_at"
	order = "desc"
	sb := strings.ToLower(sortBy)
	so := strings.ToLower(sortOrder)
	switch sb {
	case "terbaru":
		col, order = "published_at", "desc"
	case "terlama":
		col, order = "published_at", "asc"
	case "published_at", "created_at", "updated_at", "title":
		col = sb
		if so == "asc" || so == "desc" {
			order = so
		}
	}
	return col, order
}

func (r *articleRepo) List(ctx context.Context, req entity.GetArticleListRequest, satkerFilter *string) (*entity.GetArticleListResponse, error) {
	var items []Article
	db := r.db.WithContext(ctx).Model(&Article{}).Where("deleted_at IS NULL")

	if req.Q != "" {
		q := "%" + req.Q + "%"
		db = db.Where("(title ILIKE ? OR excerpt ILIKE ? OR content ILIKE ?)", q, q, q)
	}
	if req.Category != "" {
		db = db.Where("category = ?", req.Category)
	}
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
	sortBy, sortOrder := mapArticleSort(req.SortBy, req.SortOrder)
	order := sortBy + " " + sortOrder

	offset := (page - 1) * pageSize
	if err := db.Order(order + " NULLS LAST").Offset(int(offset)).Limit(int(pageSize)).Find(&items).Error; err != nil {
		return nil, err
	}

	list := make([]entity.ArticleListItem, len(items))
	for i := range items {
		list[i] = items[i].ToListItem()
	}

	totalPage := total / pageSize
	if total%pageSize > 0 {
		totalPage++
	}

	return &entity.GetArticleListResponse{
		Items:     list,
		TotalPage: totalPage,
		TotalData: total,
		Page:      page,
		PageSize:  pageSize,
	}, nil
}

func (r *articleRepo) GetByID(ctx context.Context, id string) (*entity.Article, error) {
	var a Article
	err := r.db.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&a).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return toArticleEntity(&a), nil
}

func (r *articleRepo) GetBySlug(ctx context.Context, slug string) (*entity.Article, error) {
	var a Article
	err := r.db.WithContext(ctx).Where("slug = ? AND deleted_at IS NULL", slug).First(&a).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return toArticleEntity(&a), nil
}

func toArticleEntity(a *Article) *entity.Article {
	publishedAt, createdAt, updatedAt := "", "", ""
	if a.PublishedAt != nil {
		publishedAt = a.PublishedAt.UTC().Format(time.RFC3339)
	}
	if a.CreatedAt != nil {
		createdAt = a.CreatedAt.UTC().Format(time.RFC3339)
	}
	if a.UpdatedAt != nil {
		updatedAt = a.UpdatedAt.UTC().Format(time.RFC3339)
	}
	return &entity.Article{
		ID:           a.ID,
		Title:        a.Title,
		Slug:         a.Slug,
		Content:      a.Content,
		Excerpt:      a.Excerpt,
		ThumbnailURL: a.ThumbnailURL,
		PublishedAt:  publishedAt,
		Status:       a.Status,
		AuthorName:   a.AuthorName,
		Category:     a.Category,
		SatkerID:     a.SatkerID,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}

func (r *articleRepo) SlugExists(ctx context.Context, slug, excludeID string) (bool, error) {
	var n int64
	db := r.db.WithContext(ctx).Model(&Article{}).Where("slug = ?", slug)
	if excludeID != "" {
		db = db.Where("id != ?", excludeID)
	}
	err := db.Count(&n).Error
	return n > 0, err
}

func (r *articleRepo) Create(ctx context.Context, a *entity.Article) (string, error) {
	now := time.Now().UTC()
	id := uuid.New().String()
	var publishedAt *time.Time
	if a.PublishedAt != "" {
		if t, err := time.Parse(time.RFC3339, a.PublishedAt); err == nil {
			t := t.UTC()
			publishedAt = &t
		}
	}
	dto := Article{
		ID:           id,
		Title:        a.Title,
		Slug:         a.Slug,
		Content:      a.Content,
		Excerpt:      a.Excerpt,
		ThumbnailURL: a.ThumbnailURL,
		PublishedAt:  publishedAt,
		Status:       a.Status,
		AuthorName:   a.AuthorName,
		Category:     a.Category,
		SatkerID:     a.SatkerID,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
	if err := r.db.WithContext(ctx).Create(&dto).Error; err != nil {
		return "", err
	}
	return id, nil
}

func (r *articleRepo) Update(ctx context.Context, a *entity.Article) error {
	var publishedAt *time.Time
	if a.PublishedAt != "" {
		if t, err := time.Parse(time.RFC3339, a.PublishedAt); err == nil {
			t := t.UTC()
			publishedAt = &t
		}
	}
	upd := map[string]interface{}{
		"title":        a.Title,
		"slug":         a.Slug,
		"content":      a.Content,
		"excerpt":      a.Excerpt,
		"thumbnail_url": a.ThumbnailURL,
		"published_at": publishedAt,
		"status":       a.Status,
		"author_name":  a.AuthorName,
		"category":     a.Category,
		"updated_at":   time.Now().UTC(),
	}
	return r.db.WithContext(ctx).Model(&Article{}).Where("id = ?", a.ID).Updates(upd).Error
}

func (r *articleRepo) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Article{}).Error
}
