package banksoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *bankSoalUsecase) Delete(ctx context.Context, id string, reason string) error {
	if reason == "" {
		return entity.ErrQuestionDeleteReason
	}
	q, err := u.questionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if q == nil {
		return entity.ErrQuestionNotFound
	}
	used, err := u.questionRepo.IsQuestionUsedInPackage(ctx, id)
	if err != nil {
		return err
	}
	if used {
		return entity.ErrQuestionInUseByPaket
	}
	return u.questionRepo.Delete(ctx, id, reason)
}
