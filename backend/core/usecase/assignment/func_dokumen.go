package assignment

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
)

func (u *assignmentUsecase) ListDokumenPersyaratan(ctx context.Context, jenisUjikom string) (*entity.ListDokumenPersyaratanResponse, error) {
	items, err := u.dokumenRepo.ListByJenisUjikom(ctx, jenisUjikom)
	if err != nil {
		return nil, err
	}
	return &entity.ListDokumenPersyaratanResponse{Items: items}, nil
}
