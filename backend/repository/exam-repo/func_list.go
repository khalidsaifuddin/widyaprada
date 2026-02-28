package examrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *examRepo) List(ctx context.Context, req entity.GetExamListRequest) (*entity.GetExamListResponse, error) {
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

	q := r.db.WithContext(ctx).Model(&Exam{})

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
	var rows []Exam
	orderClause := req.SortBy + " " + req.SortOrder
	if err := q.Order(orderClause).Offset(int(offset)).Limit(int(req.PageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}

	items := make([]entity.ExamListItem, len(rows))
	for i := range rows {
		e := rows[i]
		createdAt := ""
		if e.CreatedAt != nil {
			createdAt = e.CreatedAt.UTC().Format(time.RFC3339)
		}
		jadwalMulai, jadwalSelesai := "", ""
		if e.JadwalMulai != nil {
			jadwalMulai = e.JadwalMulai.UTC().Format(time.RFC3339)
		}
		if e.JadwalSelesai != nil {
			jadwalSelesai = e.JadwalSelesai.UTC().Format(time.RFC3339)
		}
		var partCount int64
		r.db.WithContext(ctx).Model(&ExamParticipant{}).Where("exam_id = ?", e.ID).Count(&partCount)
		items[i] = entity.ExamListItem{
			ID:                 e.ID,
			Code:               e.Code,
			Name:               e.Name,
			JadwalMulai:        jadwalMulai,
			JadwalSelesai:      jadwalSelesai,
			DurasiMenit:        e.DurasiMenit,
			Status:             e.Status,
			VerificationStatus: e.VerificationStatus,
			ParticipantCount:   int(partCount),
			CreatedAt:          createdAt,
		}
	}

	totalPage := total / req.PageSize
	if total%req.PageSize > 0 {
		totalPage++
	}

	return &entity.GetExamListResponse{
		Items:     items,
		TotalPage: totalPage,
		TotalData: total,
		Page:      req.Page,
		PageSize:  req.PageSize,
	}, nil
}
