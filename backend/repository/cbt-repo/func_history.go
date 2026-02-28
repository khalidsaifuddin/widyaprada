package cbtrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
)

func (r *cbtRepo) ListHistory(ctx context.Context, userID string) ([]entity.CBTHistoryItem, error) {
	var attempts []examrepo.ExamAttempt
	if err := r.db.WithContext(ctx).
		Where("user_id = ? AND submitted_at IS NOT NULL", userID).
		Order("submitted_at DESC").
		Find(&attempts).Error; err != nil {
		return nil, err
	}

	items := make([]entity.CBTHistoryItem, len(attempts))
	for i := range attempts {
		att := &attempts[i]
		startedAt := ""
		if att.StartedAt != nil {
			startedAt = att.StartedAt.UTC().Format(time.RFC3339)
		}
		submittedAt := ""
		if att.SubmittedAt != nil {
			submittedAt = att.SubmittedAt.UTC().Format(time.RFC3339)
		}
		var examName string
		r.db.WithContext(ctx).Table("exams").Select("name").Where("id = ?", att.ExamID).Scan(&examName)
		items[i] = entity.CBTHistoryItem{
			AttemptID:   att.ID,
			ExamID:      att.ExamID,
			ExamName:    examName,
			StartedAt:   startedAt,
			SubmittedAt: submittedAt,
			Score:       att.Score,
		}
	}
	return items, nil
}
