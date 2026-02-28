package banksoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *bankSoalUsecase) List(ctx context.Context, req entity.GetQuestionListRequest) (*entity.GetQuestionListResponse, error) {
	return u.questionRepo.List(ctx, req)
}
