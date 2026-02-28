package assignmentrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
)

func (r *assignmentRepo) ListMyAssignments(ctx context.Context, userID string, req entity.GetAssignmentListRequest) (*entity.GetAssignmentListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.SortBy == "" {
		req.SortBy = "jadwal_selesai"
	}
	if req.SortOrder != "asc" && req.SortOrder != "desc" {
		req.SortOrder = "asc"
	}

	// Subquery: exam_ids where user is participant
	sub := r.db.WithContext(ctx).Table("exam_participants").
		Select("exam_id").
		Where("user_id = ?", userID)

	q := r.db.WithContext(ctx).Table("exams").
		Select("exams.id, exams.name, exams.jadwal_selesai, exams.tampilkan_leaderboard").
		Where("exams.deleted_at IS NULL").
		Where("exams.id IN (?)", sub)

	// Filter status: belum_dikerjakan = no attempt/submitted, sudah_dikerjakan = has submitted attempt
	if req.Status == entity.AssignmentStatusBelumDikerjakan {
		q = q.Where("NOT EXISTS (SELECT 1 FROM exam_attempts ea WHERE ea.exam_id = exams.id AND ea.user_id = ? AND ea.submitted_at IS NOT NULL)", userID)
	} else if req.Status == entity.AssignmentStatusSudahDikerjakan {
		q = q.Where("EXISTS (SELECT 1 FROM exam_attempts ea WHERE ea.exam_id = exams.id AND ea.user_id = ? AND ea.submitted_at IS NOT NULL)", userID)
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}

	orderClause := "exams." + req.SortBy + " " + req.SortOrder
	offset := (req.Page - 1) * req.PageSize

	var rows []struct {
		ID                   string
		Name                 string
		JadwalSelesai        *time.Time
		TampilkanLeaderboard bool
	}
	if err := q.Order(orderClause).Offset(int(offset)).Limit(int(req.PageSize)).Scan(&rows).Error; err != nil {
		return nil, err
	}

	items := make([]entity.AssignmentListItem, len(rows))
	for i := range rows {
		deadline := ""
		if rows[i].JadwalSelesai != nil {
			deadline = rows[i].JadwalSelesai.UTC().Format(time.RFC3339)
		}
		var att examrepo.ExamAttempt
		err := r.db.WithContext(ctx).Where("exam_id = ? AND user_id = ?", rows[i].ID, userID).First(&att).Error
		status := entity.AssignmentStatusBelumDikerjakan
		var score *float64
		if err == nil && att.SubmittedAt != nil {
			status = entity.AssignmentStatusSudahDikerjakan
			score = att.Score
		}
		canView := rows[i].TampilkanLeaderboard && status == entity.AssignmentStatusSudahDikerjakan
		items[i] = entity.AssignmentListItem{
			ExamID:             rows[i].ID,
			ExamName:           rows[i].Name,
			Deadline:           deadline,
			Status:             status,
			Score:              score,
			CanViewLeaderboard: canView,
		}
	}

	totalPage := total / req.PageSize
	if total%req.PageSize > 0 {
		totalPage++
	}

	return &entity.GetAssignmentListResponse{
		Items:     items,
		TotalPage: totalPage,
		TotalData: total,
		Page:      req.Page,
		PageSize:  req.PageSize,
	}, nil
}
