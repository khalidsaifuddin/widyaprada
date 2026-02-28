package packagerepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *packageRepo) List(ctx context.Context, req entity.GetPackageListRequest) (*entity.GetPackageListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.SortBy == "" {
		req.SortBy = "created_at"
	}
	if req.SortOrder != "asc" && req.SortOrder != "desc" {
		req.SortOrder = "desc"
	}

	q := r.db.WithContext(ctx).Model(&QuestionPackage{})

	if req.Q != "" {
		term := "%" + req.Q + "%"
		q = q.Where("LOWER(name) LIKE LOWER(?) OR LOWER(code) LIKE LOWER(?)", term, term)
	}
	if req.Status != "" && req.Status != "all" {
		q = q.Where("status = ?", req.Status)
	}
	if req.StatusVerifikasi != "" && req.StatusVerifikasi != "all" {
		q = q.Where("verification_status = ?", req.StatusVerifikasi)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.PageSize
	var rows []QuestionPackage
	orderClause := req.SortBy + " " + req.SortOrder
	if err := q.Order(orderClause).Offset(int(offset)).Limit(int(req.PageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}

	items := make([]entity.PackageListItem, len(rows))
	for i := range rows {
		p := rows[i]
		createdAt := ""
		if p.CreatedAt != nil {
			createdAt = p.CreatedAt.UTC().Format(time.RFC3339)
		}
		var count int64
		r.db.WithContext(ctx).Model(&PackageQuestionItem{}).Where("package_id = ?", p.ID).Count(&count)
		items[i] = entity.PackageListItem{
			ID:                 p.ID,
			Code:               p.Code,
			Name:               p.Name,
			Description:        p.Description,
			Status:             p.Status,
			VerificationStatus: p.VerificationStatus,
			QuestionCount:      int(count),
			CreatedAt:          createdAt,
		}
	}

	totalPage := total / req.PageSize
	if total%req.PageSize > 0 {
		totalPage++
	}

	return &entity.GetPackageListResponse{
		Items:     items,
		TotalPage: totalPage,
		TotalData: total,
		Page:      req.Page,
		PageSize:  req.PageSize,
	}, nil
}
