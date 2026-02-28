package cbtrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
)

func (r *cbtRepo) SaveAnswer(ctx context.Context, attemptID, questionID, answerValue string) error {
	now := time.Now().UTC()

	var existing examrepo.ExamAnswer
	err := r.db.WithContext(ctx).Where("attempt_id = ? AND question_id = ?", attemptID, questionID).First(&existing).Error
	if err == nil {
		return r.db.WithContext(ctx).Model(&examrepo.ExamAnswer{}).
			Where("attempt_id = ? AND question_id = ?", attemptID, questionID).
			Updates(map[string]interface{}{"answer_value": answerValue, "updated_at": now}).Error
	}

	a := examrepo.ExamAnswer{
		AttemptID:   attemptID,
		QuestionID:  questionID,
		AnswerValue: answerValue,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}
	return r.db.WithContext(ctx).Create(&a).Error
}

func (r *cbtRepo) GetAnswers(ctx context.Context, attemptID string) (map[string]string, error) {
	var rows []examrepo.ExamAnswer
	if err := r.db.WithContext(ctx).Where("attempt_id = ?", attemptID).Find(&rows).Error; err != nil {
		return nil, err
	}
	m := make(map[string]string)
	for i := range rows {
		m[rows[i].QuestionID] = rows[i].AnswerValue
	}
	return m, nil
}
