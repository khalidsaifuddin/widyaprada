package cbtrepo

import (
	"context"
	"strings"
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

	// Hitung skor PG, MRA & B-S (Essay manual)
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
		if q.Type != entity.QuestionTypePG && q.Type != entity.QuestionTypeMRA && q.Type != entity.QuestionTypeBenarSalah {
			continue
		}
		totalWeight += q.Weight
		optVal := answers[qid]
		if optVal == "" {
			continue
		}
		if q.Type == entity.QuestionTypeMRA {
			// MRA: optVal = "id1,id2,id3", hitung dari option_weight yang benar dipilih
			selectedIDs := strings.Split(optVal, ",")
			type optInfo struct {
				ID           string
				IsCorrect    bool
				OptionWeight float64
			}
			var opts []optInfo
			r.db.WithContext(ctx).Table("question_options").Select("id, is_correct, option_weight").Where("question_id = ?", qid).Find(&opts)
			var maxPossible, earned float64
			optMap := make(map[string]optInfo)
			for _, o := range opts {
				w := o.OptionWeight
				if w <= 0 {
					w = 1
				}
				optMap[o.ID] = optInfo{ID: o.ID, IsCorrect: o.IsCorrect, OptionWeight: w}
				if o.IsCorrect {
					maxPossible += w
				}
			}
			for _, sid := range selectedIDs {
				sid = strings.TrimSpace(sid)
				if sid == "" {
					continue
				}
				o, ok := optMap[sid]
				if !ok {
					continue
				}
				if o.IsCorrect {
					earned += o.OptionWeight
				}
			}
			if maxPossible > 0 {
				earnedWeight += (earned / maxPossible) * q.Weight
			}
		} else {
			var isCorrect bool
			r.db.WithContext(ctx).Table("question_options").Select("is_correct").Where("id = ?", optVal).Scan(&isCorrect)
			if isCorrect {
				earnedWeight += q.Weight
			}
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
