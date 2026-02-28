package banksoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *bankSoalUsecase) Get(ctx context.Context, id string) (*entity.QuestionDetailResponse, error) {
	q, err := u.questionRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if q == nil {
		return nil, entity.ErrQuestionNotFound
	}
	opts, err := u.questionRepo.GetOptionsByQuestionID(ctx, id)
	if err != nil {
		return nil, err
	}
	catName := ""
	if q.CategoryID != "" {
		cat, _ := u.questionRepo.GetCategoryByID(ctx, q.CategoryID)
		if cat != nil {
			catName = cat.Name
		}
	}
	return &entity.QuestionDetailResponse{
		ID:                 q.ID,
		Code:               q.Code,
		Type:               q.Type,
		CategoryID:         q.CategoryID,
		CategoryName:       catName,
		Difficulty:         q.Difficulty,
		QuestionText:       q.QuestionText,
		AnswerKey:          q.AnswerKey,
		Weight:             q.Weight,
		Status:             q.Status,
		VerificationStatus: q.VerificationStatus,
		Options:            opts,
		CreatedAt:          q.CreatedAt,
		UpdatedAt:          q.UpdatedAt,
	}, nil
}
