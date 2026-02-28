package banksoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *bankSoalUsecase) ListCategories(ctx context.Context) ([]entity.QuestionCategory, error) {
	return u.questionRepo.ListCategories(ctx)
}
