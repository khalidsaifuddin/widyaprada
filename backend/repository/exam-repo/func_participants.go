package examrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *examRepo) GetParticipantsByExamID(ctx context.Context, examID string) ([]entity.ExamParticipantItem, error) {
	var rows []ExamParticipant
	err := r.db.WithContext(ctx).Where("exam_id = ?", examID).Find(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]entity.ExamParticipantItem, len(rows))
	for i := range rows {
		out[i] = entity.ExamParticipantItem{
			ExamID: rows[i].ExamID,
			UserID: rows[i].UserID,
		}
	}
	return out, nil
}

func (r *examRepo) SetParticipants(ctx context.Context, examID string, userIDs []string) error {
	if err := r.db.WithContext(ctx).Where("exam_id = ?", examID).Delete(&ExamParticipant{}).Error; err != nil {
		return err
	}
	if len(userIDs) == 0 {
		return nil
	}
	now := time.Now().UTC()
	rows := make([]ExamParticipant, len(userIDs))
	for i, uid := range userIDs {
		rows[i] = ExamParticipant{
			ExamID:    examID,
			UserID:    uid,
			CreatedAt: &now,
		}
	}
	return r.db.WithContext(ctx).Create(&rows).Error
}
