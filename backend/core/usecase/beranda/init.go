package beranda

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type BerandaUsecase interface {
	GetPengumuman(ctx context.Context, userID string) (*entity.BerandaPengumumanResponse, error)
}

type berandaUsecase struct {
	ujikomRepo repository.UjikomRepo
	cbtRepo    repository.CBTRepo
}

func NewBerandaUsecase(ujikomRepo repository.UjikomRepo, cbtRepo repository.CBTRepo) BerandaUsecase {
	return &berandaUsecase{
		ujikomRepo: ujikomRepo,
		cbtRepo:    cbtRepo,
	}
}
