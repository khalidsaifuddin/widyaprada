package assignmentrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
)

func (r *assignmentRepo) ListForDashboard(ctx context.Context, userID string, limit, page int64) ([]entity.DashboardAssignmentItem, int64, error) {
	if limit < 1 {
		limit = 10
	}
	if page < 1 {
		page = 1
	}

	sub := r.db.WithContext(ctx).Table("exam_participants").
		Select("exam_id").
		Where("user_id = ?", userID)

	q := r.db.WithContext(ctx).Table("exams").
		Select("exams.id, exams.name, exams.jadwal_mulai, exams.jadwal_selesai, exams.status, exams.tampilkan_leaderboard").
		Where("exams.deleted_at IS NULL").
		Where("exams.id IN (?)", sub).
		Where("exams.status IN ?", []string{entity.ExamStatusDiterbitkan, entity.ExamStatusBerlangsung, entity.ExamStatusSelesai})

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	var rows []struct {
		ID                   string
		Name                 string
		JadwalMulai          *time.Time
		JadwalSelesai        *time.Time
		Status               string
		TampilkanLeaderboard bool
	}
	if err := q.Order("exams.jadwal_selesai ASC").Offset(int(offset)).Limit(int(limit)).Scan(&rows).Error; err != nil {
		return nil, 0, err
	}

	now := time.Now().UTC()
	items := make([]entity.DashboardAssignmentItem, len(rows))
	for i := range rows {
		deadline := ""
		if rows[i].JadwalSelesai != nil {
			deadline = rows[i].JadwalSelesai.UTC().Format(time.RFC3339)
		}

		var att examrepo.ExamAttempt
		err := r.db.WithContext(ctx).Where("exam_id = ? AND user_id = ?", rows[i].ID, userID).First(&att).Error
		submitted := err == nil && att.SubmittedAt != nil

		status := entity.AssignmentStatusBelumDikerjakan
		var score *float64
		if submitted {
			status = entity.AssignmentStatusSudahDikerjakan
			score = att.Score
		}

		canStart := false
		if !submitted && rows[i].Status == entity.ExamStatusDiterbitkan && rows[i].JadwalMulai != nil && rows[i].JadwalSelesai != nil {
			canStart = !now.Before(*rows[i].JadwalMulai) && !now.After(*rows[i].JadwalSelesai)
		}

		canViewLeaderboard := rows[i].TampilkanLeaderboard && submitted

		items[i] = entity.DashboardAssignmentItem{
			ID:                 rows[i].ID,
			ExamName:           rows[i].Name,
			Deadline:           deadline,
			Status:             status,
			Score:              score,
			CanStart:           canStart,
			CanViewResult:      submitted,
			CanViewLeaderboard: canViewLeaderboard,
		}
	}
	return items, total, nil
}
