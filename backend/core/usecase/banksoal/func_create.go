package banksoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/google/uuid"
)

func (u *bankSoalUsecase) Create(ctx context.Context, req entity.CreateQuestionRequest) (*entity.QuestionDetailResponse, error) {
	existing, err := u.questionRepo.GetByCode(ctx, req.Code, "")
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, entity.ErrQuestionCodeExists
	}

	if (req.Type == entity.QuestionTypePG || req.Type == entity.QuestionTypeMRA || req.Type == entity.QuestionTypeBenarSalah) &&
		(len(req.Options) == 0 || req.AnswerKey == "") {
		return nil, entity.ErrQuestionOptionsRequired
	}

	q := &entity.QuestionDetail{
		ID:                 uuid.New().String(),
		Code:               req.Code,
		Type:               req.Type,
		CategoryID:         req.CategoryID,
		Difficulty:         req.Difficulty,
		QuestionText:       req.QuestionText,
		AnswerKey:          req.AnswerKey,
		Weight:             req.Weight,
		Status:             req.Status,
		VerificationStatus: entity.QuestionVerifBelum,
	}
	if q.Status == "" {
		q.Status = entity.QuestionStatusDraft
	}
	if q.Weight == 0 {
		q.Weight = 1
	}

	if err := u.questionRepo.Create(ctx, q); err != nil {
		return nil, err
	}

	opts := make([]entity.QuestionOption, len(req.Options))
	for i := range req.Options {
		w := req.Options[i].OptionWeight
		if w <= 0 {
			w = 1
		}
		opts[i] = entity.QuestionOption{
			ID:           uuid.New().String(),
			QuestionID:   q.ID,
			OptionKey:    req.Options[i].OptionKey,
			OptionText:   req.Options[i].OptionText,
			IsCorrect:    req.Options[i].IsCorrect,
			OptionWeight: w,
		}
	}
	if len(opts) > 0 {
		if err := u.questionRepo.CreateOptions(ctx, q.ID, opts); err != nil {
			return nil, err
		}
	}

	return u.Get(ctx, q.ID)
}
