package exam

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *examUsecase) Verify(ctx context.Context, id string) error {
	e, err := u.examRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if e == nil {
		return entity.ErrExamNotFound
	}
	e.VerificationStatus = entity.QuestionVerifSudah
	return u.examRepo.Update(ctx, e)
}

func (u *examUsecase) Unverify(ctx context.Context, id string) error {
	e, err := u.examRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if e == nil {
		return entity.ErrExamNotFound
	}
	e.VerificationStatus = entity.QuestionVerifBelum
	return u.examRepo.Update(ctx, e)
}
