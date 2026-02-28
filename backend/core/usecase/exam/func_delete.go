package exam

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *examUsecase) Delete(ctx context.Context, id string, reason string) error {
	if reason == "" {
		return entity.ErrExamDeleteReason
	}
	e, err := u.examRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if e == nil {
		return entity.ErrExamNotFound
	}
	return u.examRepo.Delete(ctx, id, reason)
}
