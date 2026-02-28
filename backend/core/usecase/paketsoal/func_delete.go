package paketsoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *paketSoalUsecase) Delete(ctx context.Context, id string, reason string) error {
	if reason == "" {
		return entity.ErrPackageDeleteReason
	}
	pkg, err := u.packageRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if pkg == nil {
		return entity.ErrPackageNotFound
	}
	used, err := u.packageRepo.IsPackageUsedInExam(ctx, id)
	if err != nil {
		return err
	}
	if used {
		return entity.ErrPackageInUseByExam
	}
	return u.packageRepo.Delete(ctx, id, reason)
}
