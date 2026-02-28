package banksoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *bankSoalUsecase) Verify(ctx context.Context, id string) error {
	q, err := u.questionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if q == nil {
		return entity.ErrQuestionNotFound
	}
	q.VerificationStatus = entity.QuestionVerifSudah
	return u.questionRepo.Update(ctx, q)
}

func (u *bankSoalUsecase) Unverify(ctx context.Context, id string) error {
	q, err := u.questionRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if q == nil {
		return entity.ErrQuestionNotFound
	}
	q.VerificationStatus = entity.QuestionVerifBelum
	return u.questionRepo.Update(ctx, q)
}
