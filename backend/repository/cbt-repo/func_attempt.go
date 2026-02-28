package cbtrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
	"gorm.io/gorm"
)

func (r *cbtRepo) GetAttemptByID(ctx context.Context, attemptID string) (examID, userID string, submitted bool, err error) {
	var att examrepo.ExamAttempt
	if err := r.db.WithContext(ctx).Where("id = ?", attemptID).First(&att).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", "", false, nil
		}
		return "", "", false, err
	}
	return att.ExamID, att.UserID, att.SubmittedAt != nil, nil
}
