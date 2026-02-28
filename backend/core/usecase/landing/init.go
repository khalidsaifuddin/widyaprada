package landing

import (
	"context"

	"github.com/ProjectWidyaprada/backend/core/entity"
	"github.com/ProjectWidyaprada/backend/core/repository"
)

type LandingUsecase interface {
	GetHome(ctx context.Context) (*entity.LandingHomeResponse, error)
}

type landingUsecase struct {
	slideRepo   repository.SlideRepo
	articleRepo repository.ArticleRepo
	linkRepo    repository.LinkRepo
	journalRepo repository.JournalRepo
}

func NewLandingUsecase(slideRepo repository.SlideRepo, articleRepo repository.ArticleRepo, linkRepo repository.LinkRepo, journalRepo repository.JournalRepo) LandingUsecase {
	return &landingUsecase{
		slideRepo:   slideRepo,
		articleRepo: articleRepo,
		linkRepo:    linkRepo,
		journalRepo: journalRepo,
	}
}
