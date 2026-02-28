package exam

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *examUsecase) Publish(ctx context.Context, id string) error {
	e, err := u.examRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if e == nil {
		return entity.ErrExamNotFound
	}
	if e.Status != entity.ExamStatusDraft {
		return entity.ErrExamAlreadyPublished
	}
	e.Status = entity.ExamStatusDiterbitkan
	return u.examRepo.Update(ctx, e)
}
