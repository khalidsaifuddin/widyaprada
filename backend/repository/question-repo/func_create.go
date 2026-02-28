package questionrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (r *questionRepo) Create(ctx context.Context, q *entity.QuestionDetail) error {
	if q.ID == "" {
		q.ID = uuid.New().String()
	}
	if err := r.ensureCategoryExists(ctx, q.CategoryID); err != nil {
		return err
	}
	now := time.Now().UTC()
	dto := Question{}.FromEntity(*q)
	dto.CreatedAt = &now
	dto.UpdatedAt = &now
	if dto.Status == "" {
		dto.Status = entity.QuestionStatusDraft
	}
	if dto.VerificationStatus == "" {
		dto.VerificationStatus = entity.QuestionVerifBelum
	}
	if err := r.db.WithContext(ctx).Create(&dto).Error; err != nil {
		return err
	}
	q.CreatedAt = now.Format(time.RFC3339)
	q.UpdatedAt = now.Format(time.RFC3339)
	return nil
}
