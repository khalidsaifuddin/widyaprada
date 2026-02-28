package assignmentrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
	"gorm.io/gorm"
)

func (r *assignmentRepo) GetAttemptByExamUser(ctx context.Context, examID, userID string) (*entity.AssignmentResultResponse, error) {
	var exam examrepo.Exam
	if err := r.db.WithContext(ctx).Where("id = ?", examID).First(&exam).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	var att examrepo.ExamAttempt
	err := r.db.WithContext(ctx).Where("exam_id = ? AND user_id = ?", examID, userID).First(&att).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &entity.AssignmentResultResponse{
				ExamID:   examID,
				ExamName: exam.Name,
			}, nil
		}
		return nil, err
	}

	resp := &entity.AssignmentResultResponse{
		ExamID:   examID,
		ExamName: exam.Name,
		Score:    att.Score,
	}
	if att.SubmittedAt != nil {
		resp.SubmittedAt = att.SubmittedAt.UTC().Format(time.RFC3339)
	}
	return resp, nil
}
