package paketsoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *paketSoalUsecase) Verify(ctx context.Context, id string) error {
	pkg, err := u.packageRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if pkg == nil {
		return entity.ErrPackageNotFound
	}
	pkg.VerificationStatus = entity.QuestionVerifSudah
	return u.packageRepo.Update(ctx, pkg)
}

func (u *paketSoalUsecase) Unverify(ctx context.Context, id string) error {
	pkg, err := u.packageRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if pkg == nil {
		return entity.ErrPackageNotFound
	}
	pkg.VerificationStatus = entity.QuestionVerifBelum
	return u.packageRepo.Update(ctx, pkg)
}
