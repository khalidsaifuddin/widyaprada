package paketsoal

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *paketSoalUsecase) List(ctx context.Context, req entity.GetPackageListRequest) (*entity.GetPackageListResponse, error) {
	return u.packageRepo.List(ctx, req)
}
