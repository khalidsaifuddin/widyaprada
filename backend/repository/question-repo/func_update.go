package questionrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (r *questionRepo) Update(ctx context.Context, q *entity.QuestionDetail) error {
	if err := r.ensureCategoryExists(ctx, q.CategoryID); err != nil {
		return err
	}
	now := time.Now().UTC()
	dto := Question{}.FromEntity(*q)
	dto.UpdatedAt = &now
	return r.db.WithContext(ctx).Model(&Question{}).Where("id = ?", q.ID).Updates(map[string]interface{}{
		"code":                dto.Code,
		"type":                dto.Type,
		"category_id":         dto.CategoryID,
		"difficulty":          dto.Difficulty,
		"question_text":       dto.QuestionText,
		"answer_key":          dto.AnswerKey,
		"weight":              dto.Weight,
		"status":              dto.Status,
		"updated_at":          dto.UpdatedAt,
	}).Error
}
