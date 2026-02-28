package cbtrepo

import (
	"context"
	"time"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/repository/exam-repo"
)

func (r *cbtRepo) SubmitAttempt(ctx context.Context, attemptID string) (*entity.CBTSubmitResponse, error) {
	var att examrepo.ExamAttempt
	if err := r.db.WithContext(ctx).Where("id = ?", attemptID).First(&att).Error; err != nil {
		return nil, err
	}
	if att.SubmittedAt != nil {
		return &entity.CBTSubmitResponse{
			AttemptID:   attemptID,
			Score:       att.Score,
			SubmittedAt: att.SubmittedAt.UTC().Format(time.RFC3339),
		}, nil
	}

	now := time.Now().UTC()

	answers, err := r.GetAnswers(ctx, attemptID)
	if err != nil {
		return nil, err
	}

	// Hitung skor PG & B-S (Essay manual)
	var totalWeight, earnedWeight float64
	questionIDs, _ := r.GetQuestionIDsForAttempt(ctx, attemptID)
	for _, qid := range questionIDs {
		var q struct {
			Type   string
			Weight float64
		}
		if err := r.db.WithContext(ctx).Table("questions").Select("type, weight").Where("id = ?", qid).Scan(&q).Error; err != nil {
			continue
		}
		if q.Type != entity.QuestionTypePG && q.Type != entity.QuestionTypeBenarSalah {
			continue
		}
		totalWeight += q.Weight
		optVal := answers[qid]
		if optVal == "" {
			continue
		}
		var isCorrect bool
		r.db.WithContext(ctx).Table("question_options").Select("is_correct").Where("id = ?", optVal).Scan(&isCorrect)
		if isCorrect {
			earnedWeight += q.Weight
		}
	}

	var score *float64
	if totalWeight > 0 {
		s := (earnedWeight / totalWeight) * 100
		score = &s
	}

	if err := r.db.WithContext(ctx).Model(&examrepo.ExamAttempt{}).
		Where("id = ?", attemptID).
		Updates(map[string]interface{}{
			"submitted_at": now,
			"score":        score,
			"updated_at":   now,
		}).Error; err != nil {
		return nil, err
	}

	return &entity.CBTSubmitResponse{
		AttemptID:   attemptID,
		Score:       score,
		SubmittedAt: now.Format(time.RFC3339),
	}, nil
}
