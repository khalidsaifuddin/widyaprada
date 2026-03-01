package questionrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (r *questionRepo) GetOptionsByQuestionID(ctx context.Context, questionID string) ([]entity.QuestionOption, error) {
	var rows []QuestionOption
	err := r.db.WithContext(ctx).Where("question_id = ?", questionID).Find(&rows).Error
	if err != nil {
		return nil, err
	}
	out := make([]entity.QuestionOption, len(rows))
	for i := range rows {
		out[i] = rows[i].ToEntity()
	}
	return out, nil
}

func (r *questionRepo) CreateOptions(ctx context.Context, questionID string, opts []entity.QuestionOption) error {
	if len(opts) == 0 {
		return nil
	}
	now := time.Now().UTC()
	for i := range opts {
		if opts[i].ID == "" {
			opts[i].ID = uuid.New().String()
		}
		opts[i].QuestionID = questionID
	}
	rows := make([]QuestionOption, len(opts))
	for i := range opts {
		w := opts[i].OptionWeight
		if w <= 0 {
			w = 1
		}
		rows[i] = QuestionOption{
			ID:           opts[i].ID,
			QuestionID:   opts[i].QuestionID,
			OptionKey:    opts[i].OptionKey,
			OptionText:   opts[i].OptionText,
			IsCorrect:    opts[i].IsCorrect,
			OptionWeight: w,
			CreatedAt:    &now,
		}
	}
	return r.db.WithContext(ctx).Create(&rows).Error
}

func (r *questionRepo) DeleteOptionsByQuestionID(ctx context.Context, questionID string) error {
	return r.db.WithContext(ctx).Where("question_id = ?", questionID).Delete(&QuestionOption{}).Error
}
