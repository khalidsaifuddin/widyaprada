package questionrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *questionRepo) List(ctx context.Context, req entity.GetQuestionListRequest) (*entity.GetQuestionListResponse, error) {
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

	q := r.db.WithContext(ctx).Model(&Question{})

	if req.Q != "" {
		term := "%" + req.Q + "%"
		q = q.Where("LOWER(question_text) LIKE LOWER(?) OR LOWER(code) LIKE LOWER(?)", term, term)
	}
	if req.Tipe != "" {
		q = q.Where("type = ?", req.Tipe)
	}
	if req.KategoriID != "" {
		q = q.Where("category_id = ?", req.KategoriID)
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
	var rows []Question
	orderClause := req.SortBy + " " + req.SortOrder
	err := q.Order(orderClause).
		Offset(int(offset)).
		Limit(int(req.PageSize)).
		Find(&rows).Error
	if err != nil {
		return nil, err
	}

	// GORM Join might not fill CategoryName - fetch separately if needed
	items := make([]entity.QuestionListItem, len(rows))
	categoryCache := make(map[string]string)
	for i := range rows {
		q := rows[i]
		catName := ""
		if q.CategoryID != "" {
			if n, ok := categoryCache[q.CategoryID]; ok {
				catName = n
			} else {
				var cat QuestionCategory
				if r.db.WithContext(ctx).Where("id = ?", q.CategoryID).First(&cat).Error == nil {
					categoryCache[q.CategoryID] = cat.Name
					catName = cat.Name
				}
			}
		}
		createdAt := ""
		if q.CreatedAt != nil {
			createdAt = q.CreatedAt.UTC().Format("2006-01-02T15:04:05Z07:00")
		}
		items[i] = entity.QuestionListItem{
			ID:                 q.ID,
			Code:               q.Code,
			Type:               q.Type,
			CategoryID:         q.CategoryID,
			CategoryName:       catName,
			Difficulty:         q.Difficulty,
			QuestionText:       q.QuestionText,
			AnswerKey:          q.AnswerKey,
			Weight:             q.Weight,
			Status:             q.Status,
			VerificationStatus: q.VerificationStatus,
			CreatedAt:          createdAt,
		}
	}

	totalPage := total / req.PageSize
	if total%req.PageSize > 0 {
		totalPage++
	}

	return &entity.GetQuestionListResponse{
		Items:     items,
		TotalPage: totalPage,
		TotalData: total,
		Page:      req.Page,
		PageSize:  req.PageSize,
	}, nil
}
