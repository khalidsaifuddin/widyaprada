package cbtrepo

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
)

func (r *cbtRepo) GetQuestionIDsForAttempt(ctx context.Context, attemptID string) ([]string, error) {
	var rows []examrepo.ExamAttemptQuestion
	if err := r.db.WithContext(ctx).Where("attempt_id = ?", attemptID).Order("sort_order").Find(&rows).Error; err != nil {
		return nil, err
	}
	ids := make([]string, len(rows))
	for i := range rows {
		ids[i] = rows[i].QuestionID
	}
	return ids, nil
}

func (r *cbtRepo) GetQuestionForCBT(ctx context.Context, questionID string) (*entity.CBTQuestionItem, error) {
	var q struct {
		ID           string
		Type         string
		QuestionText string
		Weight       float64
	}
	if err := r.db.WithContext(ctx).Table("questions").Select("id, type, question_text, weight").
		Where("id = ?", questionID).Scan(&q).Error; err != nil {
		return nil, err
	}

	item := &entity.CBTQuestionItem{
		QuestionID:   q.ID,
		Type:         q.Type,
		QuestionText: q.QuestionText,
		Weight:       q.Weight,
	}

	if q.Type == "PG" || q.Type == "BENAR_SALAH" {
		var opts []struct {
			ID         string
			OptionKey  string
			OptionText string
		}
		r.db.WithContext(ctx).Table("question_options").
			Select("id, option_key, option_text").
			Where("question_id = ?", questionID).
			Find(&opts)
		item.Options = make([]entity.CBTQuestionOption, len(opts))
		for i := range opts {
			item.Options[i] = entity.CBTQuestionOption{
				ID:         opts[i].ID,
				OptionKey:  opts[i].OptionKey,
				OptionText: opts[i].OptionText,
			}
		}
	}

	return item, nil
}
