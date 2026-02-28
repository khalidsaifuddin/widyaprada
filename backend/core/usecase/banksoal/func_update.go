package banksoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *bankSoalUsecase) Update(ctx context.Context, id string, req entity.UpdateQuestionRequest) (*entity.QuestionDetailResponse, error) {
	q, err := u.questionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if q == nil {
		return nil, entity.ErrQuestionNotFound
	}

	if req.Code != "" {
		existing, _ := u.questionRepo.GetByCode(ctx, req.Code, id)
		if existing != nil {
			return nil, entity.ErrQuestionCodeExists
		}
		q.Code = req.Code
	}
	if req.Type != "" {
		q.Type = req.Type
	}
	if req.CategoryID != "" || (req.CategoryID == "" && q.CategoryID != "") {
		q.CategoryID = req.CategoryID
	}
	if req.Difficulty != "" {
		q.Difficulty = req.Difficulty
	}
	if req.QuestionText != "" {
		q.QuestionText = req.QuestionText
	}
	if req.AnswerKey != "" {
		q.AnswerKey = req.AnswerKey
	}
	if req.Weight > 0 {
		q.Weight = req.Weight
	}
	if req.Status != "" {
		q.Status = req.Status
	}

	if (q.Type == entity.QuestionTypePG || q.Type == entity.QuestionTypeBenarSalah) &&
		len(req.Options) > 0 {
		if q.AnswerKey == "" {
			return nil, entity.ErrQuestionOptionsRequired
		}
		// Replace options
		if err := u.questionRepo.DeleteOptionsByQuestionID(ctx, id); err != nil {
			return nil, err
		}
		opts := make([]entity.QuestionOption, len(req.Options))
		for i := range req.Options {
			opts[i] = entity.QuestionOption{
				QuestionID: id,
				OptionKey:  req.Options[i].OptionKey,
				OptionText: req.Options[i].OptionText,
				IsCorrect:  req.Options[i].IsCorrect,
			}
		}
		if err := u.questionRepo.CreateOptions(ctx, id, opts); err != nil {
			return nil, err
		}
	}

	if err := u.questionRepo.Update(ctx, q); err != nil {
		return nil, err
	}
	return u.Get(ctx, id)
}
